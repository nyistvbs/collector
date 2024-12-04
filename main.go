package main

import (
	"collector/service"
	"flag"
)

// 启动服务爬取
// 启动消费者 判断端口是否存在，不存在则启动
// 消费者入库
func main() {
	flag.Parse()

	// 初始化服务
	svc := service.New()

	// 执行脚本
	svc.StartJob()

	// 启动服务
	svc.Run()
}
