package service

import (
	"collector/dao"
	"collector/queue"
	"flag"
	"log"

	"github.com/go-rod/rod"
	"github.com/pkg/errors"
)

var svc *Service

type Service struct {
	pathCfg    string
	queueCfg   string
	workersCfg int
	dao        *dao.Dao
	queue      *queue.Queue
	rod        *rod.Browser
}

func init() {
	svc = &Service{}
	flag.StringVar(&svc.pathCfg, "data", "", "")
	flag.StringVar(&svc.queueCfg, "queue", "", "")
	flag.IntVar(&svc.workersCfg, "workers", 10, "")
}

func New() *Service {
	svc.dao = dao.New()
	svc.queue = queue.New()
	svc.rod = rod.New().MustConnect()
	return svc
}

func (s *Service) StartJob() {
	// 执行消费脚本
	go s.tourism()
	// 执行爬虫脚本
	go s.crawl()
}

func (s *Service) Run() {
	s.tcp()
}

func (s *Service) Close() {
	if s.rod != nil {
		err := s.rod.Close()
		if err != nil {
			log.Fatal("rod.Close() error:", errors.WithStack(err))
		}
	}
	if s.dao != nil {
		s.dao.Close()
	}

}
