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
    //time.Sleep(time.Millisecond * 500)
    return atomic.AddInt64(&count, 1), nil
}

func main() {
    g := singleflight.Group{}

    wg := sync.WaitGroup{}

    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func(j int) {
            defer wg.Done()
            val, err, shared := g.Do("cache_key_name", a)
            if err != nil {
                fmt.Println(err)
                return
            }
            fmt.Printf("index: %d, val: %d, shard: %v\n", j, val, shared)
        }(i)
    }

    wg.Wait()
}