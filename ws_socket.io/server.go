package main

import (
    "github.com/googollee/go-socket.io"
    "github.com/googollee/go-socket.io/engineio"
    "github.com/googollee/go-socket.io/engineio/transport"
    "github.com/googollee/go-socket.io/engineio/transport/websocket"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "golang_practice_everyday/ws_socket.io/session"
    "log"
    "net/http"
    "time"
)

// go-socket.io server（客户端和服务端版本兼容问题：go-socket.io@v1.6.2仅支持客户端socket.io@v2.4.0）
func main() {
    ws := websocket.Default
    ws.CheckOrigin = func(r *http.Request) bool {
        return true
    }
    server := socketio.NewServer(&engineio.Options{
        Transports:         []transport.Transport{ws},
        SessionIDGenerator: new(session.UUIDGenerator),
    })

    server.OnConnect("/", func(s socketio.Conn) error {
        log.Println("socket.io client connected:session_id=", s.ID())
        return nil
    })

    server.OnDisconnect("/", func(s socketio.Conn, reason string) {
        log.Println("socket.io client closed", reason)
    })

    server.OnError("/", func(s socketio.Conn, e error) {
        log.Println("socket.io client meet error:", e)
    })

    registerEvent(server)

    go func() {
        defer recover()
        _ = server.Serve()
    }()
    defer func() {
        _ = server.Close()
    }()

    e := echo.New()
    e.HideBanner = true
    e.Use(middleware.Recover())
    e.Use(middleware.CORS())
    e.Any("/socket.io/", func(context echo.Context) error {
        server.ServeHTTP(context.Response(), context.Request())
        return nil
    })
    e.Logger.Fatal(e.Start(":8000"))
}

// 处理客户端发来的事件
func registerEvent(server *socketio.Server) {
    server.OnEvent("/", "message", func(s socketio.Conn, msg string) {
        log.Println("message:", msg)
        s.Emit("reply_message", "server_reply@"+msg) // 推送事件给客户端
    })

    // 收到客户端事件后，直接返回响应值给客户端该事件
    server.OnEvent("/", "ack", func(s socketio.Conn, msg string) string {
        log.Println("ack:", msg)
        return "server get ack signal, ack=" + time.Now().Format("2006-01-02 15:04:05")
    })

    // 客户端通知服务端关闭socket连接
    server.OnEvent("/", "close", func(s socketio.Conn) string {
        log.Println("close")
        // 客户端通知服务端关闭socket连接，后端可在此处理用户数据，然后通知客户端可以关闭连接了（由客户端主动关闭连接）
        return "close ok"
    })

    // namespace 是同一个服务端socket多路复用的体现
    server.OnEvent("/chat", "chat_msg", func(s socketio.Conn, msg string) {
        log.Println("chat/msg:"+msg)
        s.Emit("reply_message", "server_reply@chat@"+msg)
    })
}