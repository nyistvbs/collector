package service

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func (s *Service) web() {
	ln, err := net.Listen("tcp", ":"+s.queueCfg) // 监听 8080 端口
	if err != nil {
		log.Fatal("Error starting server: ", err)
		os.Exit(1)
	}
	defer func() {
		ln.Close()
		s.Close() // TODO 后续监听退出优雅关闭
	}()
	log.Println("Server listening on port: ", s.queueCfg)

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

	// 接收消息长度头部（假设长度为 4 字节）
	header := make([]byte, 4)
	_, err := io.ReadFull(conn, header)
	if err != nil {
		log.Fatal("Error reading header:", err)
	}
	// 假设头部是消息的长度
	messageLength := int(header[0])<<24 | int(header[1])<<16 | int(header[2])<<8 | int(header[3])

	// 接收完整的消息
	message := make([]byte, messageLength)
	_, err = io.ReadFull(conn, message)
	if err != nil {
		log.Fatal("Error reading message:", err)
	}

	// 向客户端发送响应
	_, err = conn.Write([]byte("ok"))
	if err != nil {
		log.Println("Error writing to connection:", err)
	}

	// 入队列
	s.queue.Enqueue(message)
}

func (s *Service) req(message []byte) {
	// 连接到 TCP 服务器 (假设服务器地址为 localhost:8080)
	conn, err := net.Dial("tcp", "localhost:"+s.queueCfg)
	if err != nil {
		log.Fatal("Error connecting to server:", err)
	}
	defer conn.Close()

	// 设置超时：如果连接或读取响应超过 5 秒，则报错
	conn.SetDeadline(time.Now().Add(5 * time.Second))

	// 创建一个缓冲区
	var buffer bytes.Buffer

	// 写入消息长度（4 字节）
	messageLength := uint32(len(message))
	err = binary.Write(&buffer, binary.BigEndian, messageLength)
	if err != nil {
		log.Fatal("Error writing message length:", err)
	}

	// 写入消息内容
	_, err = buffer.Write(message)
	if err != nil {
		log.Fatal("Error writing message content:", err)
	}

	_, err = conn.Write(buffer.Bytes())
	if err != nil {
		log.Fatal("Error sending message:", err)
	}

	// 接收服务器的响应
	buf := make([]byte, 1024) // 1024 字节缓冲区
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal("Error reading response:", err)
	}
	fmt.Println("发送成功，已回复:", string(buf[:n]))
}
