package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var verbose = flag.Bool("v", false, "show verbose progress messages")

// 限制目录并发数的计数信号量
var sema = make(chan struct{}, 20)

// 递归遍历以dir为根目录的整个文件树
// 并在fileSizes上发送每个已找到的文件的大小
func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subDir := filepath.Join(dir, entry.Name())
			n.Add(1)
			go walkDir(subDir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// 返回dir目录中的条目
func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}  // 获取令牌
	defer func() { <-sema }() // 释放令牌

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

// go run main.go -v /usr/local/Cellar
func main() {
	// 确定初始目录
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// 遍历文件树 单goroutine
	fileSizes := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, fileSizes)
	}
	go func() {
		n.Wait()
		close(fileSizes)
	}()

	// 定时输出结果
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(10 * time.Millisecond)
	}

	// 输出文件数和总大小
	var nfiles, nbytes int64
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes关闭,跳出到loop外(类似于goto)
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files   %.1f GB\n", nfiles, float64(nbytes)/1e9)
}
