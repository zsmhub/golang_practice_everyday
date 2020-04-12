package main

import (
	"fmt"
	"sync"
	"time"
)

// 使用 chan + sync.WaitGroup 控制 goroutine 并发数

var wg = sync.WaitGroup{}

func main() {
	userCount := 10
	ch := make(chan bool, 2) // 最多同时运行2个goroutine
	for i := 0; i < userCount; i++ {
		wg.Add(1)
		go Read(ch, i)
	}

	wg.Wait()
}

func Read(ch chan bool, i int) {
	defer wg.Done()

	ch <- true
	fmt.Printf("go func: %d, time: %d\n", i, time.Now().Unix())
	time.Sleep(time.Second)
	<-ch
}
