package main

import "fmt"

func main() {
	fmt.Println("sum: ", sum())
	fmt.Println("sum: ", sum(1, 2, 3, 4, 5))

	// 实参存在于一个slice中的时候调用变长函数
	valus := []int{1, 2, 3, 4}
	fmt.Println("sum: ", sum(valus...))

	fmt.Println("min: ", min(1, 5, 8, 0, -1))
	fmt.Println("max: ", max(1, 5, 8, 0, -1))
}

// 可传入0到n个int. vals是一个[]int
func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

// ex5.15
func min(vals ...int) int {
	if len(vals) == 0 {
		return 0
	}
	min := vals[0]
	for i := 1; i < len(vals); i++ {
		if vals[i] < min {
			min = vals[i]
		}
	}
	return min
}

// ex5.15
func max(vals ...int) int {
	if len(vals) == 0 {
		return 0
	}
	max := vals[0]
	for i := 1; i < len(vals); i++ {
		if vals[i] > max {
			max = vals[i]
		}
	}
	return max
}
