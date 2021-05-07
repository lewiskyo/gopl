package main

import "fmt"

func main() {
	months := [...]string{1: "Jan", 2: "Fe", 3: "Mar", 4: "Apr", 5: "May", 6: "June", 7: "July", 8: "Aug", 9: "Sep", 10: "Oct", 11: "Nov", 12: "Dec"}

	Q2 := months[4:7] // 前闭后开区间 4,5,6
	summer := months[6:9]
	fmt.Println(Q2)     // 	[Apr May June]
	fmt.Println(summer) //  [June July Aug]

	endlessSummer := summer[:5] // 在slice容量范围内扩展了slice
	fmt.Println(summer)
	fmt.Println(endlessSummer) // [June July Aug Sep Oct]

	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])  //  临时创建一个slice,引用a的全部元素
	fmt.Println(a) // [5 4 3 2 1 0]

	q := make([]int, 10) // 使用make构造一个 len 和 cap相同的slice
	fmt.Printf("%d %d\n", len(q), cap(q))

	p := make([]int, 10, 20) // 使用make构造一个 制定len 和 cap的slice
	fmt.Printf("%d %d\n", len(p), cap(p))
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}