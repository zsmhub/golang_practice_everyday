package main

import (
    "fmt"
    "sort"
)

type Pair struct {
    Key string
    Value int
}

type PairList []Pair

func (p PairList) Len() int {
    return len(p)
}

func (p PairList) Swap (i, j int) {
    p[i], p[j] = p[j], p[i]
}

func (p PairList) Less(i, j int) bool {
    return p[i].Value < p[j].Value
}

func sortMapByValue(m map[string]int) PairList {
    p := make(PairList, len(m))
    i := 0
    for k, v := range m {
        p[i] = Pair{k, v}
        i++
    }
    sort.Sort(p) // 此处注意：必须写 Len/Swap/Less 这三个方法
    return p
}

// 输入任意（用户，成绩）序列，可以获得成绩从低到高排序
func main() {
    m := map[string]int{"jack": 70, "peter": 96, "Tom": 70, "smith": 67}
    fmt.Println(sortMapByValue(m))
}