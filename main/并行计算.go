package main

import (
    "fmt"
    "sync"
    "time"
)

//题目
//如果我们列出10以下所有能够被3或者5整除的自然数，那么我们得到的是3，5，6和9。这四个数的和是23。
//那么请计算1000以下（不包括1000）的所有能够被3或者5整除的自然数的和。
//
//这个题目的一个思路就是：
//
//(1) 先计算1000以下所有能够被3整除的整数的和A，
//(2) 然后计算1000以下所有能够被5整除的整数和B，
//(3) 然后再计算1000以下所有能够被3和5整除的整数和C，
//(4) 使用A+B-C就得到了最后的结果。

func main() {
    var sum sync.Map
    wg := sync.WaitGroup{}

    limit := 1000 // 最大数
    startTime := time.Now()

    divider3 := 3
    divider5 := 5
    divider15 := 15

    wg.Add(3)
    go get_sum_of_divisible(limit, divider3, &sum, &wg)
    go get_sum_of_divisible(limit, divider5, &sum, &wg)
    go get_sum_of_divisible(limit, divider15, &sum, &wg)
    wg.Wait()  // 此处阻塞，等所有协程跑完才会执行下面的代码

    sum3, ok := sum.Load(divider3)
    if !ok {
        fmt.Println("数据异常，没获取到key:", divider3)
        return
    }

    sum5, ok := sum.Load(divider5)
    if !ok {
        fmt.Println("数据异常，没获取到key:", divider5)
        return
    }

    sum15, ok := sum.Load(divider15)
    if !ok {
        fmt.Println("数据异常，没获取到key:", divider15)
        return
    }

    total := sum3.(int) + sum5.(int) - sum15.(int)

    endTime := time.Now()
    fmt.Println("耗时：", endTime.Sub(startTime))
    fmt.Println("计算结果：", total)
}

func get_sum_of_divisible(num int, divider int, sum *sync.Map, wg *sync.WaitGroup) {
    defer wg.Done()
    total := 0
    for value := 0; value < num; value++ {
        if value%divider == 0 {
            total += value
        }
    }
    sum.Store(divider, total)
}
