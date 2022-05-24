package main

import (
    "errors"
    "fmt"
    "golang.org/x/sync/errgroup"
    "time"
)

func main() {
    g := errgroup.Group{}
    result := make([]error, 3) // 保存成功或者失败的结果

    g.Go(func() error {
        time.Sleep(1*time.Second)
        fmt.Println("exec #1")
        result[0] = nil
        return nil
    })

    g.Go(func() error {
        time.Sleep(3*time.Second)
        fmt.Println("exec #2")
        result[1] = errors.New("failed to exec #2")
        return result[1]
    })

    g.Go(func() error {
        time.Sleep(5 * time.Second)
        fmt.Println("exec #3")
        result[2] = errors.New("failed to exec #3")
        return result[2]
    })

    if err := g.Wait(); err != nil {
        fmt.Printf("failed: %v\n", result)
    } else {
        fmt.Println("successfully exec all")
    }
}
