package main

import (
    "flag"
    "fmt"
)

/**
 * flag 设置命令行选项
 * go run flag.go -h
 * go run flag.go -s hello
 * go run flag.go -i 1 -s hello -b true
 * go run flag.go -i 1 -b true -s hello // 这里发现个问题：-b 只能放在最后一个选项，否则会获取不到值
 */

var (
    i    *int
    b   *bool
    s *string
)

func init() {
    i = flag.Int("i", 0, "int flag value")
    b = flag.Bool("b", false, "bool flag value")
    s = flag.String("s", "default", "string flag value")
}

func main() {
    flag.Parse()

    fmt.Println("int flag:", *i)
    fmt.Println("bool flag:", *b)
    fmt.Println("s flag:", *s)
}
