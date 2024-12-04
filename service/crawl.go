package service

import (
	"collector/helper"
	"collector/model"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-rod/rod/lib/proto"
	"golang.org/x/sync/errgroup"
)

func (s *Service) crawl() {
	content, err := os.ReadFile(s.pathCfg)
	if err != nil {
		log.Fatal(err)
	}

	var data model.FileTasks
	_ = json.Unmarshal(content, &data)

	var ewg errgroup.Group
	ewg.SetLimit(s.workersCfg)

	for _, v := range data.Tasks {
		v := v
		ewg.Go(func() error {
			s.crawlQueue(v)
			return nil
		})
	}
	_ = ewg.Wait()
}

func (s *Service) crawlQueue(item *model.TaskItem) {
	fmt.Println(item.Url)

	// 打开目标页面
	page := s.rod.MustPage(item.Url).MustWaitLoad()

	_ = page.SetUserAgent(&proto.NetworkSetUserAgentOverride{
		UserAgent: item.Headers.UserAgent,
	})

	// TODO
	time.Sleep(5 * time.Second)
	//page.MustWait("span")

	// 获取页面 HTML
	html, err := page.HTML()
	if err != nil {
		log.Fatal("Error getting HTML:", err)
	}

	// 解析 HTML
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatal(err)
	}

	// 提取数据
	//fmt.Println(doc.Find("data-testid").Text())
	data := make([]*model.CrawlData, 0)
	doc.Find("div").Each(func(i int, str *goquery.Selection) {
		key, exist := str.Attr("data-testid")
		if !exist || !helper.InStringArray([]string{"listing-card-title", "price-availability-row", "listing-card-subtitle"}, key) {
			return
		}

		fmt.Println("Title:", str.Text())
		data = append(data, &model.CrawlData{
			Key: key,
			Val: str.Text(),
		})

		//酒店名称 listing-card-title
		//明星
		//价格 price-availability-row
		//税前价格
		//入住日期 listing-card-subtitle
		//退房日期
		//客人
	})

	buf, err := json.Marshal(data)
	if err != nil {
		log.Fatal("xxx")
		return
	}
	s.req(buf)
}
