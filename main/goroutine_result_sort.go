package main

import (
    "fmt"
    "time"
)

// 按序返回多个 goroutine 的结果集

func run(task_id int, sleeptime int, ch chan string) {

    time.Sleep(time.Duration(sleeptime) * time.Second)
    ch <- fmt.Sprintf("task id %d , sleep %d second", task_id, sleeptime)
    return
}

func main() {
    input := []int{3, 2, 1}
    chs := make([]chan string, len(input)) // 定义一级 channel
    startTime := time.Now()
    fmt.Println("Multirun start")
    for i, sleeptime := range input {
        chs[i] = make(chan string) // 定义二级 channel
        go run(i, sleeptime, chs[i])
    }

    for _, ch := range chs {
        fmt.Println(<-ch)
    }

    endTime := time.Now()
    fmt.Printf("Multissh finished. Process time %s. Number of tasks is %d", endTime.Sub(startTime), len(input))
}
