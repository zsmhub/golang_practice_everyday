package main

import (
    "fmt"
    "sort"
)

func main() {
    m := map[int]string{1: "a", 0: "b", 4: "d", 2: "c"}

    var keys []int
    for k := range m {
        keys = append(keys, k)
    }
    sort.Ints(keys)

    for _, k := range keys {
        fmt.Println("key:", k, "value:", m[k])
    }
}
