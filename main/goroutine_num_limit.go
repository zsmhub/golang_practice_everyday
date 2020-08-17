package main

import (
    "fmt"
    "sync"
    "time"
)

// 使用 chan + sync.WaitGroup 限制 goroutine 并发数
// 为什么要限制协程数量? golang的go关键字并发实在是太简单，但是带来的问题是由于硬件和网络状况的限制，
// 不受控制的增加协程是非常危险的做法，甚至有可能搞垮数据库之类的应用!

func main() {
    var wg = sync.WaitGroup{} // 用于保证所有的协程都能跑完『在实际项目中可能不用，因为主进程一直处于挂起状态』

    userCount := 20 // 总共需要执行的协程数量

    ch := make(chan struct{}, 3) // 最多同时运行 n 个 goroutine
    for i := 0; i < userCount; i++ {
        wg.Add(1)
        ch <- struct{}{}
        go func(i int) {
            defer wg.Done()

            fmt.Printf("go func: %d, time: %d\n", i, time.Now().Unix())
            time.Sleep(time.Second)
            <-ch
        }(i)
    }

    wg.Wait()
}
