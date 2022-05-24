package main

import (
    "context"
    "golang.org/x/sync/semaphore"
    "log"
    "sync"
    "time"
)

// 限制 goroutine 并发数方案
// 为什么要限制协程数量? golang的go关键字并发实在是太简单，但是带来的问题是由于硬件和网络状况的限制，
// 不受控制的增加协程是非常危险的做法，甚至有可能搞垮数据库之类的应用!

// 一、使用 chan + sync.WaitGroup 限制 goroutine 并发数
func main1() {
    var (
        total = 20                     // 总共需要执行的协程数量
        wg    = sync.WaitGroup{}       // 用于保证所有的协程都能跑完『在实际项目中可能不用，因为主进程一直处于挂起状态』
        ch    = make(chan struct{}, 3) // 最多同时运行 n 个 goroutine
    )

    for i := 0; i < total; i++ {
        wg.Add(1)
        ch <- struct{}{}
        go func(i int) {
            defer wg.Done()

            log.Printf("go func: %d, time: %d\n", i, time.Now().Unix())
            time.Sleep(time.Second)
            <-ch
        }(i)
    }

    wg.Wait()
}

// 二、使用信号量限制 goroutine 并发数
func main() {
    var (
        semaWeight int64 = 3  // 最多同时运行 n 个 goroutine
        total            = 20 // 总共需要执行的协程数量
        sema             = semaphore.NewWeighted(semaWeight)
        ctx              = context.Background()
    )

    for i := 0; i < total; i++ {
        if err := sema.Acquire(ctx, 1); err != nil {
            log.Printf("sema acquire err: %v\n", err)
            continue
        }

        go func(i int) {
            log.Printf("go func: %d, time: %d\n", i, time.Now().Unix())
            time.Sleep(time.Second)
            sema.Release(1)
        }(i)
    }

    // 请求所有的worker，这样能确保前面的worker都执行完
    if err := sema.Acquire(ctx, semaWeight); err != nil{
        log.Printf("sema acquire err: %v\n", err)
    }
}
