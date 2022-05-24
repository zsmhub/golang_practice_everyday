package main

import (
    "bufio"
    "fmt"
    "go.etcd.io/etcd/client/v3"
    "go.etcd.io/etcd/client/v3/concurrency"
    "go.etcd.io/etcd/client/v3/experimental/recipes"
    "log"
    "math/rand"
    "os"
    "strings"
    "time"
)

// etcd实现分布式读写锁
var (
    etcdAddrRw = "http://127.0.0.1:23791,http://127.0.0.1:23792,http://127.0.0.1:23793"
    rwLockName = "my-test-lock"
)

func main() {
    rand.Seed(time.Now().UnixNano())

    // 解析etcd地址
    endpoints := strings.Split(etcdAddrRw, ",")

    // 创建etcd的client
    cli, err := clientv3.New(clientv3.Config{Endpoints: endpoints})
    if err != nil {
        log.Fatal(err)
    }

    defer cli.Close()

    // 创建session
    s1, err := concurrency.NewSession(cli) // 节点宕机对应 session 销毁，持有的锁会被释放
    if err != nil {
        log.Fatal(err)
    }

    defer s1.Close()

    m1 := recipe.NewRWMutex(s1, rwLockName)

    // 从命令行读取命令
    consolescanner := bufio.NewScanner(os.Stdin)
    log.Println("请输入指令w/r：")
    for consolescanner.Scan() {
        action := consolescanner.Text()
        switch action {
        case "w": // 请求写锁
            testWriteLocker(m1)
        case "r": // 请求读锁
            testReadLocker(m1)
        default:
            fmt.Println("unknown action")
        }
        log.Println("请输入指令w/r：")
    }
}

func testWriteLocker(m1 *recipe.RWMutex) {
    // 请求写锁
    log.Println("acquiring write lock")
    if err := m1.Lock(); err != nil {
        log.Fatal(err)
    }

    log.Println("acquired write lock")

    // 等待一段时间
    time.Sleep(time.Duration(rand.Intn(10)+3) * time.Second)

    // 释放写锁
    if err := m1.Unlock(); err != nil {
        log.Fatal(err)
    }

    log.Println("released write lock")
}

func testReadLocker(m1 *recipe.RWMutex) {
    // 请求读锁
    log.Println("acquiring read lock")
    if err := m1.RLock(); err != nil {
        log.Fatal(err)
    }

    log.Println("acquired read lock")

    // 等待一段时间
    time.Sleep(time.Duration(rand.Intn(10)+3) * time.Second)

    // 释放写锁
    if err := m1.RUnlock(); err != nil {
        log.Fatal(err)
    }

    log.Println("released read lock")
}
