package main

import (
	"fmt"
	"net"
)

// 简单C/S模型通讯
func main() {
	// 主动发起连接请求
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("Dial err:", err)
		return
	}

	defer conn.Close()

	// 发送数据
	_, err = conn.Write([]byte("Are u ready?"))
	if err != nil {
		fmt.Println("Write err:", err)
		return
	}
}
