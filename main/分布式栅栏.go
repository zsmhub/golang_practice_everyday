package main

import (
    "bufio"
    "fmt"
    "go.etcd.io/etcd/client/v3"
    recipe "go.etcd.io/etcd/client/v3/experimental/recipes"
    "log"
    "os"
    "strings"
)

// 分布式栅栏
var (
    etcdAddrB   = "http://127.0.0.1:23791,http://127.0.0.1:23792,http://127.0.0.1:23793"
    barrierName = "my-test-queue"
)

func main() {

    // 解析etcd地址
    endpoints := strings.Split(etcdAddrB, ",")

    // 创建etcd的client
    cli, err := clientv3.New(clientv3.Config{Endpoints: endpoints})
    if err != nil {
        log.Fatal(err)
    }
    defer cli.Close()

    // 创建/获取栅栏
    b := recipe.NewBarrier(cli, barrierName)

    // 从命令行读取命令
    fmt.Println("请输入指令hold/release/wait：")
    consolescanner := bufio.NewScanner(os.Stdin)
    for consolescanner.Scan() {
        action := consolescanner.Text()
        items := strings.Split(action, " ")
        switch items[0] {
        case "hold": // 持有这个barrier
            if err := b.Hold(); err != nil {
                fmt.Printf("hold fail: %v\n", err)
            } else {
                fmt.Println("hold success")
            }
        case "release": // 释放这个barrier
            if err := b.Release(); err != nil {
                fmt.Printf("released fail: %v\n", err)
            } else {
                fmt.Println("released success")
            }
        case "wait": // 等待barrier被释放
            if err := b.Wait(); err != nil {
                fmt.Printf("wait fail: %v\n", err)
            } else {
                fmt.Println("after wait")
            }
        case "quit", "exit": // 退出
            return
        default:
            fmt.Println("unknown action")
        }
        fmt.Println("请输入指令hold/release/wait：")
    }
}
