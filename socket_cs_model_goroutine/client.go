package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// 并发C/S模型通讯
func main() {
	// 主动发起连接请求
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("Dial err:", err)
		return
	}

	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("->")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		// 发送数据到服务端
		_, err = conn.Write([]byte(text))
		if err != nil {
			fmt.Println("Write err:", err)
			continue
		}

		// 获取服务端的返回值
		buf := make([]byte, 2048)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Read err:", err)
			break
		}
		fmt.Println("server response:", string(buf[:n]))

		if "exit" == text {
			break
		}
	}
}
