package main

import (
    "fmt"
    "github.com/gorilla/websocket"
    "log"
    "net/http"
)

// 简单 websocket 服务端
var upgrader = websocket.Upgrader{}

func main() {
    http.HandleFunc("/socket", socketHandler)
    http.HandleFunc("/", home)
    log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func socketHandler(w http.ResponseWriter, r *http.Request) {
    server, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println("Error during connection upgradation: ", err)
        return
    }

    defer server.Close()

    for {
        messageType, message, err := server.ReadMessage()
        if err != nil {
            log.Println("Error: ", err)
            break
        }

        log.Println("Server Received: ", string(message))
        err = server.WriteMessage(messageType, message)
        if err != nil {
            log.Println("Error: ", err)
            break
        }
    }
}

func home(w http.ResponseWriter, r *http.Request) {
    if _, err := fmt.Fprintf(w, "Index page"); err != nil {
        log.Println(err)
    }
}
