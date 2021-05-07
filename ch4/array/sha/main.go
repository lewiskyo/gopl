package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)

	fmt.Println("diffCount: ", diffCount(c1, c2))

	zero(&c1)
	fmt.Println(c1)
}

// ex4.1 统计SHA256散列中不同的位数
// 数组作为参数传递,是值传递,传递的是一个副本.
func diffCount(lhs, rhs [32]uint8) (count int) {
	for i := 0; i < 32; i++ {
		if lhs[i] != rhs[i] {
			count++
		}
	}
	return count
}

// 数组清零
func zero(prt *[32]uint8) {
	*prt = [32]uint8{}
}
