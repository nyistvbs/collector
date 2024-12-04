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

//
//func b() {
//	content, err := os.ReadFile(file)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	var data model.FileTasks
//	_ = json.Unmarshal(content, &data)
//
//	for _, v := range data.Tasks {
//		fmt.Println(v.Url)
//		a(v)
//	}
//}
//
//func a(item *model.TaskItem) {
//	// 创建浏览器实例
//	browser := rod.New().MustConnect()
//	defer browser.MustClose()
//
//	// 打开目标页面
//	page := browser.MustPage(item.Url).MustWaitLoad()
//
//	_ = page.SetUserAgent(&proto.NetworkSetUserAgentOverride{
//		UserAgent: item.Headers.UserAgent,
//	})
//	//err := page.
//	//if err != nil {
//	//	return
//	//}
//
//	// TODO
//	time.Sleep(5 * time.Second)
//	//page.MustWait("span")
//
//	// 获取页面 HTML
//	html, err := page.HTML()
//	if err != nil {
//		log.Fatal("Error getting HTML:", err)
//	}
//
//	// 解析 HTML
//	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// 提取数据
//	//fmt.Println(doc.Find("data-testid").Text())
//	doc.Find("div").Each(func(i int, s *goquery.Selection) {
//		val, exist := s.Attr("data-testid")
//		if !exist || val != "listing-card-title" {
//			return
//		}
//		fmt.Println("Title:", s.Text())
//		//酒店名称 listing-card-title
//		//明星
//		//价格 price-availability-row
//		//税前价格
//		//入住日期 listing-card-subtitle
//		//退房日期
//		//客人
//
//	})
//
//	//// 打印 HTML 内容
//	//fmt.Println(html)
//
//}
//
//
//
//func c() {
//	q := &Queue{items: list.New()}
//
//	// 入队操作
//	q.Enqueue(10)
//	q.Enqueue(20)
//	q.Enqueue(30)
//
//	fmt.Println("Queue length:", q.Length()) // 输出 3
//
//	// 出队操作
//	item, err := q.Dequeue()
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println("Dequeued item:", item) // 输出 10
//
//	// 查看队列长度
//	fmt.Println("Queue length after dequeue:", q.Length()) // 输出 2
//
//	// 判断队列是否为空
//	fmt.Println("Is queue empty?", q.IsEmpty()) // 输出 false
//}
//
//func handleConnection(conn net.Conn) {
//	defer conn.Close()
//
//	// 读取客户端数据
//	buf := make([]byte, 1024)
//	_, err := conn.Read(buf)
//	if err != nil {
//		log.Println("Error reading from connection:", err)
//		return
//	}
//
//	// 打印收到的消息
//	fmt.Printf("Received message: %s\n", string(buf))
//
//	// 向客户端发送响应
//	_, err = conn.Write([]byte("Hello from server!"))
//	if err != nil {
//		log.Println("Error writing to connection:", err)
//	}
//}
//
//func d() {
//	// 监听 TCP 连接
//	ln, err := net.Listen("tcp", ":8080") // 监听 8080 端口
//	if err != nil {
//		log.Fatal("Error starting server: ", err)
//		os.Exit(1)
//	}
//	defer ln.Close()
//	fmt.Println("Server listening on port 8080...")
//
//	for {
//		// 等待并接收客户端连接
//		conn, err := ln.Accept()
//		if err != nil {
//			log.Println("Error accepting connection: ", err)
//			continue
//		}
//
//		// 处理连接
//		go handleConnection(conn)
//	}
//}
