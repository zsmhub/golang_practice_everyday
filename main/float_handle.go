package main

import (
	"fmt"
)

// 注意Golang浮点型的默认舍入规则是四舍六入五成双
// 四舍六入五成双是一种比较精确比较科学的计数保留法，是一种数字修约规则，又名银行家舍入法。它比通常用的四舍五入法更加精确。
func main() {
	fmt.Printf("9.8249	=>	%0.2f(四舍)\n", 9.8249)
	fmt.Printf("9.82671	=>	%0.2f(六入)\n", 9.82671)
	fmt.Printf("9.8351	=>	%0.2f(五后非零就进一)\n", 9.8351)
	fmt.Printf("9.82501	=>	%0.2f(五后非零就进一)\n", 9.82501)
	fmt.Printf("9.8250	=>	%0.2f(五后为零看奇偶，五前为偶应舍去)\n", 9.8250) // 这里的规则要特别注意
	fmt.Printf("9.8350	=>	%0.2f(五后为零看奇偶，五前为奇要进一)\n", 9.8350)
}
