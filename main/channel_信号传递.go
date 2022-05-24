package main

import (
    "fmt"
    "time"
)

// 题目：有 4 个 goroutine，编号为 1、2、3、4。每秒钟会有一个 goroutine 打印出它自己的编号，要求你编写程序，让输出的编号总是按照 1、2、3、4、1、2、3、4……这个顺序打印出来。
type Token struct{} // 令牌

func main() {
    workers := []chan Token{make(chan Token), make(chan Token), make(chan Token), make(chan Token)}

    for k := range workers {
        go newWorker(k, workers[k], workers[(k+1)%4])
    }

    workers[0] <- Token{}

    select {}
}

func newWorker(id int, currentChan chan Token, nextChan chan Token) {
    for {
        ch := <-currentChan // 获取令牌
        fmt.Println(id + 1)
        time.Sleep(time.Second)
        nextChan <- ch // 令牌传递
    }
}
