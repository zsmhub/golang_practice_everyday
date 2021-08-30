package main

import (
    "fmt"
    "net"
)

// TCP 粘包问题复现
func main() {
    // 主动发起连接请求
    conn, err := net.Dial("tcp", "127.0.0.1:8000")
    if err != nil {
        fmt.Println("Dial err:", err)
        return
    }

    defer conn.Close()

    for i := 0; i < 100; i++ {

        msg := `Hello world!`

        _, err = conn.Write([]byte(msg))
        if err != nil {
            fmt.Println("Write err:", err)
            break
        }
        // time.Sleep(time.Microsecond*100) // 发包间隔时间变长，就不会出现粘包
    }
}
