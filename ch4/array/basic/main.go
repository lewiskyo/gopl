package main

import "fmt"

func main() {
	var a [3]int
	fmt.Println(a[0])
	fmt.Println(len(a) - 1)

	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	var q = [3]int{1, 2, 3}
	fmt.Println(q[2])
	var p = [3]int{1, 2}
	fmt.Println(p[2]) // default "0"

	z := [...]int{1, 2, 3}
	fmt.Printf("%T\n", z) // "[3]int"

	r := [...]int{99:-1} // 100个元素,下标0~98的值均为0
	fmt.Println(r[0])  // default "0"
}
