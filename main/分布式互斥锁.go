package main

import (
    "context"
    "go.etcd.io/etcd/client/v3"
    "go.etcd.io/etcd/client/v3/concurrency"
    "log"
    "math/rand"
    "strings"
    "time"
)

// etcd实现分布式互斥锁
var (
    addr     = "http://127.0.0.1:23791,http://127.0.0.1:23792,http://127.0.0.1:23793"
    lockName = "my-test-lock"
)

func main() {
    rand.Seed(time.Now().UnixNano())

    // etcd地址
    endpoints := strings.Split(addr, ",")

    // 生成一个etcd client
    cli, err := clientv3.New(clientv3.Config{Endpoints: endpoints})
    if err != nil {
        log.Fatal(err)
    }
    defer cli.Close()

    useLock(cli) // 测试锁
}

func useLock(cli *clientv3.Client) {
    // 为锁生成session「节点宕机对应 session 销毁，持有的锁会被释放」
    s1, err := concurrency.NewSession(cli)
    if err != nil {
        log.Fatal(err)
    }
    defer s1.Close()

    // 得到一个分布式锁
    locker := concurrency.NewMutex(s1, lockName)

    // 请求锁
    log.Println("acquiring lock")
    if err := locker.Lock(context.TODO()); err != nil {
        log.Fatal(err)
    }
    log.Println("acquired lock")

    // 等待一段时间
    time.Sleep(time.Duration(rand.Intn(30)+3) * time.Second)

    // 释放锁
    if err := locker.Unlock(context.TODO()); err != nil {
        log.Fatal(err)
    }

    log.Println("released lock")
}
