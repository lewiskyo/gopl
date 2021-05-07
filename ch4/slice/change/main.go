package main

import "fmt"

func main() {
	s := []string{"1", "", "3"}
	s = noempty(s)
	fmt.Println(s)

	q := []string{"1", "", "3"}
	q = noempty2(q)
	fmt.Println(q)

	stack := make([]int, 0)
	stack = stackPush(stack, 1)
	stack = stackPush(stack, 2)
	stack = stackPush(stack, 3)
	stack = stackPush(stack, 4)
	stack = stackPush(stack, 5)
	fmt.Println("stackTop: ", stackTop(stack))
	stack = stackPop(stack)
	fmt.Println("stackTop: ", stackTop(stack))

	fmt.Println("stack: ", stack)
	stack = remove(stack, 2)
	fmt.Println("after delete stack: ", stack)

	ustrings := []string{"haha","xixi","xixi","xixi","q","w","w","w","e"}
	ustrings = (unique(ustrings))
	fmt.Println(ustrings)
}

// 将数组切片中,空字符串元素去掉
func noempty(s []string) []string {
	newLen := 0
	for i := 0; i < len(s); i++ {
		if s[i] != "" {
			s[newLen] = s[i]
			newLen++
		}
	}
	return s[:newLen]
}

// 删除数组相邻相同元素 ex4.5
func unique(s []string) []string {
	compareIdx := 0
	replaceIdx := 1
	for i := 1; i < len(s); i++ {
		if s[i] != s[compareIdx] {
			s[replaceIdx] = s[i]
			replaceIdx++
			compareIdx = i
		}
	}
	return s[:replaceIdx]
}

// 版本2,使用append
func noempty2(s []string) []string {
	out := s[:0] //  引用原始slice的新的零长度的slice
	for _, elem := range s {
		if elem != "" {
			out = append(out, elem)
		}
	}
	return out
}

// 入栈
func stackPush(stack []int, elem int) []int {
	return append(stack, elem)
}

// 获取栈顶元素
func stackTop(stack []int) int {
	if len(stack) > 0 {
		return stack[len(stack)-1]
	}
	return 0
}

// 出栈(删除切片最后一个元素)
func stackPop(stack []int) []int {
	if len(stack) > 0 {
		return stack[:len(stack)-1]
	}
	return stack
}

// 删除切片中指定下标元素
func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
