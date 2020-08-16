package main

import (
    "fmt"
)

// 切片注意事项：两个切片共享同一个底层数组。如果一个切片修改了该底层数组的共享部分，另一个切片也能感知到，运行下面的代码：
func main() {
    sliceInit := []int{10, 20, 30, 40, 50} // 长度和容量均为 5

    // 子切片的容量为底层数组的长度减去切片在底层数组的开始偏移量，即新切片容量为 5-1=4
    newSlice := sliceInit[1:3] // 新切片：长度为 2，容量为 4
    fmt.Println(newSlice)

    sliceInit[1] = 200
    fmt.Println(newSlice)

    // 当切片容量不足的时候，Go 会以原始切片容量的 2 倍建立新的切片
    slice := make([]int, 2, 10)
    slice1 := slice[1:2]
    slice2 := append(slice1, 1)
    slice2 = append(slice2, 1)
    //slice2 = append(slice2, 1)
    //slice2 = append(slice2, 1)
    //slice2 = append(slice2, 1)
    //slice2 = append(slice2, 1)
    //slice2 = append(slice2, 1)
    //slice2 = append(slice2, 1)
    //slice2 = append(slice2, 1)
    //slice2 = append(slice2, 1)
    //slice2 = append(slice2, 1)
    //slice2 = append(slice2, 1)
    slice2[0] = 10001
    fmt.Println(slice)
    fmt.Println(slice1)
    fmt.Println(slice2)
    fmt.Println(cap(slice2)) // slice 容量
    fmt.Println(len(slice2)) // slice 长度

    // 切片深拷贝示例
    test()
}

// 切片深拷贝示例
func test() {
    a := make([]int, 10)
    for i := 0; i < 10; i++ {
        a[i] = i
    }
    b := a[1:4]
    b[1] = 22 // 浅拷贝：值变化会影响到切片a

    var c = make([]int, 3)
    copy(c, b)
    c[1] = 222 // 深拷贝: 值变化不会影响到切片a、b

    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)
}
