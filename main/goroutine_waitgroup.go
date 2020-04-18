package main

import (
	"fmt"
	"sync"
	"time"
)

// 练习题：实现10个人并发读数据，其中一人可以获得锁去读数据库的数据（读完并更新一/二级缓存），其他人直接读取永久缓存。【减少用户等待时间】

func main() {
	fmt.Printf("开始时间: %s\n", time.Now().Format("2006-01-02 15:04:05"))

	num := 10
	wg := sync.WaitGroup{} // 使用sync.WaitGroup方式保证所有goroutine都运行完毕
	wg.Add(num)
	for i := 0; i < num; i++ {
		condition := true
		if i == 0 {
			condition = false
		}
		go func(c bool) {
			defer wg.Done()

			if c {
				ret := foreverCached()
				fmt.Println(ret)
			} else {
				ret := getDbData()
				fmt.Println(ret)
			}
		}(condition)
	}

	wg.Wait()

	fmt.Printf("结束时间: %s\n", time.Now().Format("2006-01-02 15:04:05"))
}

// 永久缓存
func foreverCached() string {
	ret := "永久缓存"
	return ret
}

// 从数据库读数据
func getDbData() string {
	time.Sleep(time.Second * 3)
	ret := "从数据库获取数据"
	return ret
}
