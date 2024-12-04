package service

import (
	"collector/model"
	"context"
	"encoding/json"
	"fmt"
	"log"
)

func (s *Service) tourism() {
	for {
		if s.queue.IsEmpty() {
			continue
		}

		val, err := s.queue.Dequeue()
		if err != nil {
			continue
		}

		ctx := context.Background()

		data := make([]*model.CrawlData, 0)
		err = json.Unmarshal(val.([]byte), &data)
		if err != nil {
			log.Fatal("")
			continue
		}

		// 处理逻辑
		rows := &model.TourismDB{}
		for _, v := range data {
			fmt.Println(v)
			switch v.Key {
			case "listing-card-title":
				rows.HotelName = v.Val
			case "price-availability-row":
				rows.Price = v.Val
			}
		}

		//酒店名称 listing-card-title
		//明星
		//价格 price-availability-row
		//税前价格
		//入住日期 listing-card-subtitle
		//退房日期
		//客人

		buf, _ := json.Marshal(rows)
		fmt.Println(string(buf))
		if rows.HotelName != "" {
			s.dao.TourismInsert(ctx, rows)
		}
	}

}
