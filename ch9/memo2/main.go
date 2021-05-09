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

type Memo struct {
	f     Func
	mu    sync.Mutex // 保护cache
	cache map[string]result
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
		cache: make(map[string]result),
	}
}

// 非并发安全
func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mu.Lock()
	res, ok := memo.cache[key]
	if !ok {
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
	}
	memo.mu.Unlock()
	return res.value, res.err
}

func main() {
	m := New(httpGetBody)

	var wg sync.WaitGroup
	incomingUrls := []string{"https://golang.org", "https://golang.org", "http://baidu.com", "https://golang.org", "https://play.golang.org", "http://www.gopl.io"}
	for _, url := range incomingUrls {
		wg.Add(1)
		// 开辟多个go同时访问
		go func(url string) { // 需要参数,不能捕获迭代变量
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
