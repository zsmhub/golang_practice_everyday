package main

import (
    "fmt"
    "net"
    "strings"
)

// 并发C/S模型通讯
func main() {
    // 创建监听
    listener, err := net.Listen("tcp", ":50000") // tcp 不能使用大写
    if err != nil {
        fmt.Println("listen err:", err)
        return
    }

    defer listener.Close() // 主进程结束时，关闭listener

    fmt.Println("服务器等待客户端建立连接...")

    for {
        // 等待客户端连接请求
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("accept err:", err)
            break
        }

        go HandleConn(conn)
    }
}

// 处理用户请求
func HandleConn(conn net.Conn) {
    defer conn.Close() // 使用结束，断开与开户段连接

    // 获取客户端的网络地址信息
    addr := conn.RemoteAddr().String()
    fmt.Println(addr, " connect successful")

    for {
        buf := make([]byte, 2048) // 创建2k大小的缓冲区，用于read

        // 读取用户数据
        n, err := conn.Read(buf)
        if err != nil {
            fmt.Println("read err:", err)
            break
        }

        fmt.Printf("[%s]: %s\n", addr, string(buf[:n]))

        // 把数据转换为大写，再发送给客户端
        conn.Write([]byte(strings.ToUpper(string(buf[:n]))))

        // 连接中断
        if "exit" == string(buf[:n]) {
            fmt.Println(addr, "exit")
            break
        }
    }
}
