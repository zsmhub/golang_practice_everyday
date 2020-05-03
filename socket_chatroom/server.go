package main

import (
    "fmt"
    "net"
    "strings"
    "time"
)

// 并发聊天室-服务端

// 定义用户结构体类型
type Client struct {
    Messages chan string // 发送消息的通道
    Name     string      // 用户名
    Addr     string      // 网络地址：ip+port
}

// 定义全局 map 存储在线用户 key:ip+port, value: Client
var onlineMap = make(map[string]Client)

// 定义全局 channel 处理消息【带缓冲，可以提高消息处理效率】
var message = make(chan string, 20)

// 推送消息到客户端
func WriteMsgToClient(clnt Client, conn net.Conn) {
    // 循环跟踪 clnt.Messages, 有消息则读走，Write 给客户端
    for msg := range clnt.Messages {
        conn.Write([]byte(msg))
    }
}

func MakeMsg(clnt Client, msg string) string {
    buf := "[" + clnt.Addr + "]" + clnt.Name + ": " + msg
    return buf
}

// 处理客户端连接请求
func HandleConnect(conn net.Conn) {
    defer conn.Close()

    // 获取新链接上来的用户的网络地址（ip+port）
    netAddr := conn.RemoteAddr().String()
    fmt.Println("netAddr:", netAddr)

    // 给新用户创建结构体，用户名、网络地址一样
    clnt := Client{make(chan string), netAddr, netAddr}

    // 将新创建的结构体，添加到 map 中， key 值为获取到的网络地址（ip+port）
    onlineMap[netAddr] = clnt

    // 新创建一个 goroutine，专门给当前客户端发送消息
    go WriteMsgToClient(clnt, conn)

    // 广播新用户上线
    message <- MakeMsg(clnt, "login")

    hasData := make(chan bool) // 检测用户是否有消息发送

    // 循环读取用户发送的消息，广播给在线用户
    go func() {
        buf := make([]byte, 2048) // 存储读到的用户信息
        for {
            n, err := conn.Read(buf)
            if n == 0 {
                continue
            }
            if err != nil {
                fmt.Println("Read err:", err)
                return
            }

            msg := string(buf[:n])

            // 处理消息
            if msg == "exit" {
                fmt.Printf("用户%s退出登录\n", clnt.Name)

                close(clnt.Messages)
                delete(onlineMap, netAddr)
                message <- MakeMsg(clnt, "logout")
                conn.Write([]byte(msg)) // 返回退出信息给客户端
                break
            } else if msg == "who" { // 如果用户发送了“who”指令，则展示所有在线用户
                conn.Write([]byte("user list:\n"))
                for _, user := range onlineMap {
                    userInfo := user.Addr + ":" + user.Name + "\n"
                    conn.Write([]byte(userInfo)) // 写给当前用户
                }
            } else if len(msg) >= 8 && msg[:6] == "rename" { // 修改用户名: rename|...
                newName := strings.Split(msg, "|")[1]
                clnt.Name = newName
                onlineMap[netAddr] = clnt
                conn.Write([]byte("rename successful\n"))
            } else {
                message <- MakeMsg(clnt, msg)
            }

            hasData <- true
        }
    }()

    // 超时处理，60秒内没发言则自动退出聊天室
    for {
        select {
        case <-hasData:
        case <-time.After(60 * time.Second):
            delete(onlineMap, netAddr)
            message <- MakeMsg(clnt, "time out leave")
            conn.Write([]byte("timeout")) // 通知当前用户断开连接
            return                        // 结束当前应用
        }
    }
}

// 用户消息广播
func Manager() {
    // 循环读取 message 通道中的数据
    for {
        // 通道 message 中有数据读到 msg 中，没有则阻塞
        msg := <-message

        // 一旦执行到这里，说明 message 中有数据了，解除阻塞，遍历 map
        for _, clnt := range onlineMap {
            clnt.Messages <- msg // 把从 Message 通道中读到的数据，写到 client 的 C 通道中
        }
    }
}

func main() {
    fmt.Println("server start...")

    // 创建监听 socket
    listener, err := net.Listen("tcp", "127.0.0.1:8000")
    if err != nil {
        fmt.Println("Listen err:", err)
        return
    }
    defer listener.Close()

    // 创建 goroutine 处理消息
    go Manager()

    // 循环接受客户端连接请求
    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Accept err:", err)
            continue // 失败，监听其他客户端连接请求
        }
        fmt.Println("有新客户端连接进来...")

        // 给新连接的客户端，单独创建一个 goroutine，处理客户端连接请求
        go HandleConnect(conn)
    }
}
