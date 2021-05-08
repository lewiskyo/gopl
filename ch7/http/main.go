package main

import (
	"fmt"
	"log"
	"net/http"
)

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

// 查询商品列表
func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

// 查询某个商品价格
func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func main() {
	db := database{"shoes": 50, "socks": 5}
	// http.HandleFunc注册db.list 和 db.price时, 都会转换为 HandlerFunc类型, type HandlerFunc func(ResponseWriter, *Request)
	// 而HandleFunc也实现了Handler接口 中 serveHTTP方法
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	// http://127.0.0.1:8080/lists
	// http://127.0.0.1:8080/price?item=socks
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
