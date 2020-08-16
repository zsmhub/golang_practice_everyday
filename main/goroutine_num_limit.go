package main

import (
    "fmt"
    "sync"
    "time"
)

// 使用 chan + sync.WaitGroup 限制 goroutine 并发数
// 为什么要限制协程数量? golang的go关键字并发实在是太简单，但是带来的问题是由于硬件和网络状况的限制，
// 不受控制的增加协程是非常危险的做法，甚至有可能搞垮数据库之类的应用!

var wg = sync.WaitGroup{}

func main() {
    userCount := 10

    limitNum := 2
    ch := make(chan struct{}, limitNum) // 最多同时运行2个goroutine

    for i := 0; i < userCount; i++ {
        wg.Add(1)
        go Read(ch, i)

        // 闭包方式
        //go func(i int) {
        //    defer wg.Done()
        //
        //    ch <- struct{}{}
        //    fmt.Printf("go func: %d, time: %d\n", i, time.Now().Unix())
        //    time.Sleep(time.Second)
        //    <-ch
        //}(i)
    }

    wg.Wait()
}

func Read(ch chan struct{}, i int) {
    defer wg.Done()

    ch <- struct{}{}
    fmt.Printf("go func: %d, time: %d\n", i, time.Now().Unix())
    time.Sleep(time.Second)
    <-ch
}
