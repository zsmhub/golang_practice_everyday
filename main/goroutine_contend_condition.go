package main

import (
    "fmt"
    "sync"
    "time"
)

// 多服务器模式-使用redis控制共享变量

// 单服务器模式-方案一：使用互斥锁或读写锁来处理goroutine的竞争条件
func main() {
    startTime := time.Now().UnixNano()

    nums := 0
    total := 0

    wg := sync.WaitGroup{}
    wg.Add(100000)

    var mu sync.Mutex

    for i := 1; i <= 100000; i++ {
        nums += i

        go func(i int) {
            mu.Lock()
            total += i
            mu.Unlock()
            wg.Done()
        }(i)
    }

    wg.Wait()

    fmt.Printf("total:%d sum %d\n", total, nums)

    endTime := time.Now().UnixNano()
    fmt.Println("耗时(纳秒)：", endTime-startTime)
}

// 单服务器模式-方案二：相比方案一，效率低
/*package main

import (
	"fmt"
	"sync"
)

// 如何在程序中避免竞争条件呢？
// 方案：【采用channel方式】避免从多个 goroutine 访问同一个变量，例如创建一个唯一能够访问该变量的 goroutine，
// 从而将这个变量限制在单个 goroutine 内部，其他 goroutine 通过`通道`来受限的发送查询或变更变量的请求
// 以下是银行存款管理案例

var num int = 100
var deposits = make(chan int) // 存款
var balances = make(chan int) // 金额

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance: // 亮点：等到上面的deposits通道清空后再跑此处[main主程序结束->for循环结束]
		}
	}
}

func main() {
	go teller()

	var wg sync.WaitGroup
	wg.Add(num)

	for i := 0; i < num; i++ {
		go func(amount int) {
			deposits <- amount
			wg.Done()
		}(100)
	}

	wg.Wait()

	fmt.Printf("balance = %d\n", <-balances)
}*/
