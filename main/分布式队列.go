
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

// etcd实现分布式队列
var (
    etcdAddrQ     = "http://127.0.0.1:23791,http://127.0.0.1:23792,http://127.0.0.1:23793"
    queueName = "my-test-queue"
)

func main() {
    // 解析etcd地址
    endpoints := strings.Split(etcdAddrQ, ",")

    // 创建etcd的client
    cli, err := clientv3.New(clientv3.Config{Endpoints: endpoints})
    if err != nil {
        log.Fatal(err)
    }
    defer cli.Close()

    // 创建/获取队列
    q := recipe.NewQueue(cli, queueName)

    // 从命令行读取命令
    fmt.Println("请输入指令：push/pop，多参数用空格间隔")
    consolescanner := bufio.NewScanner(os.Stdin)
    for consolescanner.Scan() {
        action := consolescanner.Text()
        items := strings.Split(action, " ")
        switch items[0] {
        case "push": // 加入队列
            if len(items) != 2 {
                fmt.Println("must set value to push")
                continue
            }
            q.Enqueue(items[1]) // 入队
        case "pop": // 从队列弹出
            v, err := q.Dequeue() // 出队
            if err != nil {
                log.Fatal(err)
            }
            fmt.Println(v) // 输出出队的元素
        case "quit", "exit": //退出
            return
        default:
            fmt.Println("unknown action")
        }

        fmt.Println("请输入指令：push/pop，多参数用空格间隔")
    }
}