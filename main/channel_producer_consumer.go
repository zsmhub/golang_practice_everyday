package main

import "fmt"

// 单项channel及应用：生产者消费者模式
func main() {
    ch := make(chan int) // 创建一个双向channel

    // 新建一个goroutine， 模拟生产者，产生数据，写入 channel
    go producer(ch) // channel传参， 传递的是引用。

    // 主go程，模拟消费者，从channel读数据，打印到屏幕
    consumer(ch) // 与 producer 传递的是同一个 channel
}

// 此通道只能写，不能读。
func producer(out chan<- int) {
    for i := 0; i < 10; i++ {
        out <- i * i // 将 i*i 结果写入到只写channel
    }
    close(out)
}

// 此通道只能读，不能写
func consumer(in <-chan int) {
    for num := range in { // 从只读channel中获取数据
        fmt.Println("num =", num)
    }
}
