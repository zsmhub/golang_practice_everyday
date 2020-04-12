package main

import "fmt"

// 如何在程序中避免竞争条件呢？
// 方案：【采用channel方式】避免从多个 goroutine 访问同一个变量，例如创建一个唯一能够访问该变量的 goroutine，
// 从而将这个变量限制在单个 goroutine 内部，其他 goroutine 通过`通道`来受限的发送查询或变更变量的请求
// 以下是银行存款管理案例

var num int = 100
var deposits = make(chan int) // 存款
var balances = make(chan int) // 金额
var sigovers = make(chan struct{}, num)

func Deposit(amount int) {
	deposits <- amount
	sigovers <- struct{}{}
}

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
	for i := 0; i < num; i++ {
		go Deposit(100)
		<-sigovers
	}
	fmt.Printf("balance = %d\n", <-balances)
}
