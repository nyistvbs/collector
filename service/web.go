package service

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func (s *Service) tcp() {
	ln, err := net.Listen("tcp", ":8080") // 监听 8080 端口
	if err != nil {
		log.Fatal("Error starting server: ", err)
		os.Exit(1)
	}
	defer ln.Close()
	fmt.Println("Server listening on port 8080...")

	for {
		// 等待并接收客户端连接
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Error accepting connection: ", err)
			continue
		}

		// 处理连接
		go s.resp(conn)
	}
}

func (s *Service) resp(conn net.Conn) {
	defer conn.Close()

	// 读取客户端数据
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		log.Println("Error reading from connection:", err)
		return
	}

	// 打印收到的消息
	fmt.Printf("Received message: %s\n", string(buf))

	// 向客户端发送响应
	_, err = conn.Write([]byte("Hello from server!"))
	if err != nil {
		log.Println("Error writing to connection:", err)
	}

	// 入队列
	s.queue.Enqueue(buf)
}

func (s *Service) req(message string) {
	// 连接到 TCP 服务器 (假设服务器地址为 localhost:8080)
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("Error connecting to server:", err)
	}
	defer conn.Close()

	// 设置超时：如果连接或读取响应超过 5 秒，则报错
	conn.SetDeadline(time.Now().Add(5 * time.Second))

	// 发送请求数据
	_, err = conn.Write([]byte(message))
	if err != nil {
		log.Fatal("Error sending message:", err)
	}
	fmt.Println("Message sent:", message)

	// 接收服务器的响应
	buffer := make([]byte, 1024) // 1024 字节缓冲区
	n, err := conn.Read(buffer)
	if err != nil {
		log.Fatal("Error reading response:", err)
	}
	fmt.Println("Received from server:", string(buffer[:n]))
}
