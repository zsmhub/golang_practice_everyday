package main

import (
    "fmt"
    "golang.org/x/sync/singleflight"
    "sync"
    "sync/atomic"
)

// 防缓存击穿代码示例

var count = int64(0)

func a() (interface{}, error) {
    //time.Sleep(time.Millisecond * 500) // 去掉注释可查看高并发取值情况
    return atomic.AddInt64(&count, 1), nil
}

func main() {
    g := singleflight.Group{}

    wg := sync.WaitGroup{}

    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func(j int) {
            defer wg.Done()

            // 直接返回结果集
            val, err, shared := g.Do("cache_key_name", a)
            if err != nil {
                fmt.Println(err)
                return
            }
            fmt.Printf("index: %d, val: %d, shared: %v\n", j, val, shared)

            // 返回值为 channel 方式
            //result := g.DoChan("cache_key_name2", a)
            //ret, _ := <-result
            //fmt.Printf("channel->index: %d, val: %d, shared: %v\n", j, ret.Val, ret.Shared)
        }(i)
    }

    wg.Wait()
}