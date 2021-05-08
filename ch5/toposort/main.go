package main

import (
	"fmt"
	"sort"
)

// 反应了所有课程和先决课程的关系
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"liner algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":      {"discrete math"},
	"databases":            {"data structures"},
	"discrete math":        {"intro to programing"},
	"formal languages":     {"discrete math"},
	"networks":             {"operating systems"},
	"operating systems":    {"data structures", "computer organization"},
	"programing languages": {"data structures", "computer organization"},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)

	// 当一个匿名函数需要进行递归,必须先声明一个变量然后将匿名函数赋给这个变量;
	// 如果合成一个声明,函数字面量将不能存在于visitAll变量的作用于中,不能递归自己
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	// 如果不对key排序,visitAll得到的结果将不唯一
	sort.Strings(keys)

	visitAll(keys)
	return order
}
