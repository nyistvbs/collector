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
	table      string
	password   string
	user       string
	workersCfg int
	mode       int
	dao        *dao.Dao
	queue      *queue.Queue
	rod        *rod.Browser
}

func init() {
	svc = &Service{}
	flag.StringVar(&svc.pathCfg, "data", "", "")
	flag.StringVar(&svc.queueCfg, "host", "8080", "")
	flag.IntVar(&svc.workersCfg, "workers", 10, "")
	flag.StringVar(&svc.user, "user", "root", "")
	flag.StringVar(&svc.password, "password", "123456", "")
	flag.StringVar(&svc.table, "table", "tourism", "")
	flag.IntVar(&svc.mode, "mode", 1, "")
}

func New() *Service {
	// 初始化数据层
	svc.dao = dao.New(svc.user, svc.password, svc.table)
	// 初始化队列
	svc.queue = queue.New()
	// 初始化爬虫
	svc.rod = rod.New().MustConnect()

	return svc
}

func (s *Service) StartJob() {
	// 执行消费脚本
	if s.mode == 1 {
		go s.tourism()
	}

	// 执行爬虫脚本
	if s.mode == 2 {
		s.crawl()
	}
}

func (s *Service) Run() {
	if s.mode == 1 {
		s.web()
	}
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
