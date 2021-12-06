package main

import (
    "fmt"
    "time"
)

// goroutine 超时控制『注意 goroutine 泄漏问题』

func Run(task_id, sleeptime, timeout int, ch chan string) {
    ch_run := make(chan string, 1) // 设置缓存区，解决 goroutine 泄漏问题（没有数据传输推荐使用空结构体 struct{}，空结构体占用内存为 0）
    go logic(task_id, sleeptime, ch_run)
    select {
    case re := <-ch_run:
        ch <- re
    case <-time.After(time.Duration(timeout) * time.Second):
        re := fmt.Sprintf("task id %d , timeout", task_id)
        ch <- re
    }
}

func logic(task_id, sleeptime int, ch chan string) {
    time.Sleep(time.Duration(sleeptime) * time.Second)
    ch <- fmt.Sprintf("task id %d , sleep %d second", task_id, sleeptime) // 导致 goroutine 泄露：外层导致提前退出，由于没有接收者且无缓存区，发送者(sender)会一直阻塞，导致协程不能退出。
    return
}

func main() {
    input := []int{3, 2, 1}
    timeout := 2
    chs := make([]chan string, len(input))
    startTime := time.Now()
    fmt.Println("Multirun start")
    for i, sleeptime := range input {
        chs[i] = make(chan string)
        go Run(i, sleeptime, timeout, chs[i])
    }

    for _, ch := range chs {
        fmt.Println(<-ch)
    }
    endTime := time.Now()
    fmt.Printf("Multissh finished. Process time %s. Number of task is %d", endTime.Sub(startTime), len(input))
}
