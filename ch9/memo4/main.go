package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

// 向指定url获取页面数据方法
func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

type entry struct {
	res   result
	ready chan struct{} // res 准备好后会被关闭
}

type Memo struct {
	f     Func
	mu    sync.Mutex
	cache map[string]*entry
}

// Func是用于获取缓存的函数类型
type Func func(string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	return &Memo{
		f:     f,
		cache: make(map[string]*entry),
	}
}

func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		// 对key的第一次访问,这个goroutine负责计算数据和广播数据
		// 已准备完毕消息, 本地先构造好entry
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()

		e.res.value, e.res.err = memo.f(key)

		close(e.ready) // 计算数据完成后, 广播数据
	} else {
		// 对这个key的重复访问
		memo.mu.Unlock()
		// 等待数据准备完毕
		<-e.ready
	}
	return e.res.value, e.res.err
}

// 实现了并发,重复抑制,非阻塞缓存
func main() {
	m := New(httpGetBody)

	var wg sync.WaitGroup
	incomingUrls := []string{"https://golang.org", "http://baidu.com", "https://golang.org", "https://play.golang.org", "http://www.gopl.io"}
	for _, url := range incomingUrls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				log.Print(err)
			}
			fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte))) // value类型转换为[]byte
		}(url)
	}
	wg.Wait()
}
