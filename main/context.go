package main

import (
    "context"
    "time"
)

// 使用 context 操作 goroutine
func main() {
    // cancel
    ctx, cancel := context.WithCancel(context.Background())
    go work(ctx, "work1")
    time.Sleep(time.Second * 3)
    cancel()
    time.Sleep(time.Second)

    // timeout
    ctx2, timeCancel := context.WithTimeout(context.Background(), time.Second*3)
    go work(ctx2, "time cancel")
    time.Sleep(time.Second)
    timeCancel() // 调用此方法会提前中断 goroutine，不调用则按设置的超时时间中断 goroutine

    // deadline
    ctx3, deadlineCancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*3))
    go work(ctx3, "deadline cancel")
    time.Sleep(time.Second * 2)
    deadlineCancel() // 调用此方法会提前中断 goroutine，不调用则按设置的超时时间中断 goroutine

    time.Sleep(time.Second * 3)
}

func work(ctx context.Context, name string) {
    for {
        select {
        case <-ctx.Done():
            println(name, " get message to quit")
            return
        default:
            println(name, " is running", time.Now().String())
            time.Sleep(time.Second)
        }

    }
}
