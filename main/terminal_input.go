package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 在这个例子中，每次输入一个字符串并按下 enter 键，我们会通过 \n 这个关键字符来区分字符串，
// 如果你想对比刚才输入的字符串，我们还需要调用 replace 方法来去除掉 \n 然后再进行比较。
func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Shell")
	fmt.Println("---------------------")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n') // 注意此处使用单引号

		// 如果你想让这个程序在 Windows 系统下运行，那么你必须将代码的 text 替换为 text = strings.Replace(text,"\r\n","",-1)
		// 因为 Windowss 系统使用的行结束符和 unix 系统是不同的。
		text = strings.Replace(text, "\n", "", -1)

		if strings.Compare("hi", text) == 0 {
			fmt.Println("hello, Yourself")
		} else {
			fmt.Println("server: ", text)
		}
	}

}
