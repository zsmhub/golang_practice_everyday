package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
    "strings"
)

// 并发聊天室-客户端
func main() {
    // 主动发起连接请求
    conn, err := net.Dial("tcp", "127.0.0.1:8000")
    if err != nil {
        fmt.Println("Dial err:", err)
        return
    }

    defer conn.Close()

    reader := bufio.NewReader(os.Stdin)

    // 获取用户的输入
    go func() {
        for {
            fmt.Print("->")
            text, _ := reader.ReadString('\n')
            text = strings.Replace(text, "\n", "", -1)

            // 发送数据到服务端
            _, err = conn.Write([]byte(text))
            if err != nil {
                fmt.Println("Write err:", err)
                break
            }

            if "exit" == text {
                break
            }
        }
    }()

    // 获取服务端的返回值
    for {
        buf := make([]byte, 2048)
        n, err := conn.Read(buf)
        if n == 0 {
            continue
        }
        if err != nil {
            fmt.Println("Read err:", err)
            break
        }
        response := string(buf[:n])
        fmt.Println(response)
        if response == "timeout" || response == "exit" {
            break
        }
        fmt.Print("->")
    }
}
