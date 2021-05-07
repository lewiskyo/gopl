package main

import "fmt"

// 复合map
var graph = make(map[string]map[string]bool)

func main() {
	ages := make(map[string]int)
	ages["haha"] = 123
	ages["xixi"] = 234

	// 字面量初始化
	ages2 := map[string]int{
		"haha": 123,
		"xixi": 234,
	}
	fmt.Println(ages2["haha"])

	delete(ages, "haha")    // 移除元素 ages["haha"]
	delete(ages, "noexist") // 即使键不在map中,操作也是安全的

	ages["bob"] += 1 // 如果键对应的元素不存在,就返回值类型的零值
	fmt.Println(ages["bob"])

	// 试图获取map零值会发生编译错误,因为map的增长可能会导致已有元素被重新散列到新的存储位置,使得获取的地址失效
	// _ = &ages["bob"]

	// 迭代的顺序不是固定的,要想固定输出,先获取所有的key,然后排序
	for name, age := range ages2 {
		fmt.Printf("name: %s, age: %d\n", name, age)
	}

	if _, ok := ages2["bob"]; !ok {
		fmt.Println("bob not in ages2")
	} else {
		fmt.Println("bob in ages2")
	}

	addEdge("a","b")
	addEdge("a","c")
	addEdge("b","c")

	fmt.Println("a to d? ", hasEdge("a","d"))
	fmt.Println("a to c? ", hasEdge("a","c"))
}

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		// 延迟初始化map,key第一次出现时初始化
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	// 即使from和to都不存在,依然可以给出一个有意义的值
	return graph[from][to]
}
