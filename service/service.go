package service

import (
	"collector/dao"
	"collector/queue"
	"flag"
)

var svc *Service

type Service struct {
	pathCfg    string
	queueCfg   string
	workersCfg int
	dao        *dao.Dao
	queue      *queue.Queue
}

func init() {
	svc = &Service{}
	flag.StringVar(&svc.pathCfg, "data", "", "")
	flag.StringVar(&svc.queueCfg, "queue", "", "")
	flag.IntVar(&svc.workersCfg, "workersCfg", 0, "")
}

func New() *Service {
	svc.dao = dao.New()
	svc.queue = queue.New()
	return svc
}

func (s *Service) StartJob() {
	// 执行爬虫脚本
	s.crawl()
	// 执行消费脚本
}

func (s *Service) Run() {
	s.tcp()
}
