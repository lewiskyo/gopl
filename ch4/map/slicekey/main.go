package main

import "fmt"

func main() {
	m := make(map[string]int)
	l1 := []string{"a", "b"}
	l2 := []string{"a", "c"}

	m[k(l1)]++
	m[k(l2)]++
	m[k(l2)]++

	fmt.Println("l1: ", m[k(l1)])
	fmt.Println("l2: ", m[k(l2)])
}

// slice是不可比较类型,因此不能作为map的key
// 可以自定义方法,将字符串slice转换为string
// k函数也可以是其他方法,例如将字符串转换为小写等
func k(list []string) string {
	return fmt.Sprintf("%q", list)
}
