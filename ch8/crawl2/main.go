package main

import (
	"fmt"
	"gopl/ch8/links"
	"log"
	"os"
)

// 限制并行数最大为20
var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // 获取令牌
	list, err := links.Extract(url)
	<-tokens // 释放令牌

	if err != nil {
		log.Print(err)
	}
	return list
}

// go run main.go https://golang.org
func main() {
	// 待过滤url列表
	worklist := make(chan []string)

	var n int // 等待发送到任务列表的url数量

	// 从命令行参数开始
	n++
	go func() {
		worklist <- os.Args[1:]
	}()

	// 并发爬取web, 主goroutine负责url去重
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}
