package main

import (
	"fmt"
	"strings"
)


// 变量x在main函数中返回squares函数后依旧存在,x隐藏在函数变量f中
func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main() {
	s := strings.Map(func(r rune) rune { return r + 1 }, "HAL-9000")
	fmt.Println(s)

	f := squares()
	fmt.Println(f())  // "1"
	fmt.Println(f())  // "4"
	fmt.Println(f())  // "9"
	fmt.Println(f())  // "16"
}
