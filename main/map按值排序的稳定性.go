package main

import (
    "fmt"
    "sort"
)

type person struct {
    Name string
    Age  int
}

type personSlice []person

func (s personSlice) Len() int           { return len(s) }
func (s personSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s personSlice) Less(i, j int) bool { return s[i].Age < s[j].Age }

// sort不保证排序的稳定性（两个相同的值，排序之后相对位置不变），排序的稳定性由sort.Stable来保证。
func main() {
    a := personSlice{
        {
            Name: "AAA",
            Age:  55,
        },
        {
            Name: "BBB",
            Age:  22,
        },
        {
            Name: "CCC",
            Age:  0,
        },
        {
            Name: "DDD",
            Age:  22,
        },
        {
            Name: "EEE",
            Age:  11,
        },
    }
    sort.Stable(a)
    fmt.Println(a)
}
