package main

import (
    "github.com/gorilla/websocket"
    "io/ioutil"
    "log"
    "os"
    "os/signal"
    "time"
)

// 简单 websocket 客户端
var (
    done      chan interface{}
    interrupt chan os.Signal
)

func main() {
    done = make(chan interface{}) // 监听websocket是否关闭连接
    interrupt = make(chan os.Signal)

    signal.Notify(interrupt, os.Interrupt)

    socketUrl := "ws://127.0.0.1:8000"
    client, resp, err := websocket.DefaultDialer.Dial(socketUrl, nil)
    if err != nil {
        log.Fatal(err)
    }
    if resp != nil && resp.StatusCode != 101 {
        b, _ := ioutil.ReadAll(resp.Body)
        log.Fatalf("resp err=%s", string(b))
    }

    defer client.Close()
    go receiveHandler(client)

    for {
        select {
        case <-time.After(1 * time.Second):
            // 发消息给服务端
            err := client.WriteMessage(websocket.TextMessage, []byte("Hello World!"))
            if err != nil {
                log.Println(err)
                return
            }

        case <-interrupt:
            // 客户端中断，发送关闭连接消息给服务端
            log.Println("Received SIGINT interrupt signal. Closing all pending connections")

            err := client.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
            if err != nil {
                log.Println("Error during closing websocket:", err)
                return
            }

            select {
            case <-done:
                log.Println("Receiver Channel Closed! Exiting....")
            case <-time.After(1 * time.Second):
                log.Println("Timeout in closing receiving channel. Exiting....")
            }

            return
        }
    }
}

func receiveHandler(connection *websocket.Conn) {
    defer close(done)
    for {
        _, msg, err := connection.ReadMessage()
        if err != nil {
            log.Println(err)
            return
        }
        log.Printf("Client Received: %s\n", msg)
    }
}
