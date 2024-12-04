package service

import (
	"collector/helper"
	"collector/model"
	"encoding/json"
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
			log.Println("正在爬取URL", v.Url)
			s.crawlQueue(v)
			return nil
		})
	}
	_ = ewg.Wait()
}

func (s *Service) crawlQueue(item *model.TaskItem) {
	// 打开目标页面
	page := s.rod.MustPage(item.Url).MustWaitLoad()

	_ = page.SetUserAgent(&proto.NetworkSetUserAgentOverride{
		UserAgent: item.Headers.UserAgent,
	})

	// TODO 等待页面完全加载
	time.Sleep(5 * time.Second)

	// 获取页面 HTML
	html, err := page.HTML()
	if err != nil {
		log.Fatal("Error getting HTML:", err)
	}

	// 解析 HTML
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatal("goquery.NewDocumentFromReader err:", err)
	}

	// 提取数据
	data := make([]*model.CrawlData, 0)
	doc.Find("div").Each(func(i int, str *goquery.Selection) {
		key, exist := str.Attr("data-testid")
		if !exist || !helper.InStringArray([]string{"listing-card-title", "price-availability-row", "listing-card-subtitle"}, key) {
			return
		}

		log.Println("提取数据 标签:", key, " 内容:", str.Text())
		data = append(data, &model.CrawlData{
			Key: key,
			Val: str.Text(),
		})
	})

	buf, err := json.Marshal(data)
	if err != nil {
		log.Fatal("序列化错误")
		return
	}
	s.req(buf)
}
