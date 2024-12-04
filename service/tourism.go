package service

import (
	"collector/model"
	"context"
	"encoding/json"
	"log"
)

func (s *Service) tourism() {
	// TODO 后续采用堵塞队列
	defer func() {
		if err := recover(); err != nil {
			log.Fatal("recover:", err)
		}
	}()
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
			log.Fatal("序列化失败", err)
		}

		// 处理逻辑 TODO 业务逻辑处理爬虫数据并去重
		rows := &model.TourismDB{}
		for _, v := range data {
			switch v.Key {
			case "listing-card-title":
				rows.HotelName = v.Val
			case "price-availability-row":
				rows.Price = v.Val
			}
		}

		// 入库
		if rows.HotelName != "" {
			s.dao.TourismInsert(ctx, rows)
			log.Println("入库成功 数据:", rows)
		}
	}

}
