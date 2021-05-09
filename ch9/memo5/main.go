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

type request struct {
	key      string
	response chan<- result // 传递到memo,memo只需往通道发消息
}

type entry struct {
	res   result
	ready chan struct{} // res 准备好后会被关闭
}

func (e *entry) call(f Func, key string) {
	// 执行函数
	e.res.value, e.res.err = f(key)
	// 通知数据已经准备完毕
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	// 等待该数据准备完毕
	<-e.ready
	// 向客户端发送数据
	response <- e.res
}

type Memo struct{ requests chan request }

// Func是用于获取缓存的函数类型
type Func func(string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

// New返回f的函数记忆,客户端之后需要调用Close
func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

func (memo *Memo) Get(key string) (interface{}, error) {
	response := make(chan result)
	request := request{
		key:      key,
		response: response,
	}
	memo.requests <- request
	res := <-response
	return res.value, res.err
}

func (memo *Memo) Close() {
	close(memo.requests)
}

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for req := range memo.requests {
		e := cache[req.key]
		if e == nil {
			// 对这个key的第一次调用
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key)
		}
		go e.deliver(req.response)
	}
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
	m.Close()
}