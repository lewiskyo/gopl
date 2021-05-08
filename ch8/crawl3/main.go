package main

import (
	"fmt"
	"gopl/ch8/links"
	"log"
	"os"
)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)

	if err != nil {
		log.Print(err)
	}
	return list
}

// go run main.go https://golang.org
func main() {
	worklist := make(chan []string)  // 可能有重复的URL列表
	unseenLinks := make(chan string) // 去重后的URL列表

	// 向任务列表中添加命令行参数
	go func() {
		worklist <- os.Args[1:]
	}()

	// 创建20个爬虫goroutine来获取每个不可见链接
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				// 开辟新goroutine进行上交任务操作,避免阻塞crawl工作
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	// 主goroutine负责对URL列表进行去重
	// 并把没有爬取过的条目发送给爬虫程序
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}
