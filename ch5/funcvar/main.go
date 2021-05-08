package main

import (
	"fmt"
	"strings"
)

func square(n int) int {
	return n * n
}

func negative(n int) int {
	return -n
}

func product(m, n int) int {
	return m * n
}

// 对字符进行加1操作
func add1(r rune) rune {
	return r + 1
}

func main() {
	f := square
	fmt.Println(f(2))

	ff := negative
	fmt.Println(ff(2))

	fff := product
	fmt.Println(fff(3, 2))

	// 对每个字符进行add1操作,将结果连接起来
	fmt.Println(strings.Map(add1, "HAL-9000"))
	fmt.Println(strings.Map(add1, "VMS"))
	fmt.Println(strings.Map(add1, "Admix"))
}
