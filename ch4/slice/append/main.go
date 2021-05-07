package main

import "fmt"

func main() {
	var runes []rune // rune类型用于存储UTF-8字符串
	for _, r := range "Hello,世界" {
		runes = append(runes, r)
	}

	fmt.Printf("%q\n", runes)

	x := make([]int, 0)
	fmt.Printf("%d %d\n", len(x), cap(x))
	x = appendInt(x, 1)
	fmt.Printf("%d %d\n", len(x), cap(x))
	x = appendInt(x, 1)
	fmt.Printf("%d %d\n", len(x), cap(x))
	x = appendInt(x, 1)
	fmt.Printf("%d %d\n", len(x), cap(x))

	x = appendInt2(x, 1, 2, 3)
	fmt.Printf("%d %d\n", len(x), cap(x))
}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// slice 仍有增长空间,扩展slice内容
		z = x[:zlen]
	} else {
		// slice已无空间,为它分配一个新的底层数组
		// 为了达到分摊线性复杂性,容量扩展一倍(实际底层算法不是这样)
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // copy有返回值,返回实际上复制的元素个数,是两个slice长度的较小值,所以不存在由于元素复制而导致的索引越界问题
	}
	z[len(x)] = y
	return z
}

// y是可变长度参数
func appendInt2(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	copy(z[len(x):], y)
	return z
}
