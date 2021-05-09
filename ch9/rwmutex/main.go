package main

import (
	"fmt"
	"sync"
)

var icons map[string]string
var m sync.RWMutex

func loadIcons() {
	icons = map[string]string{
		"spades.png":   "spades.png",
		"hearts.png":   "hearts.png",
		"diamonds.png": "diamonds.png",
		"clubs.png":    "clubs.png",
	}
}

func icon(name string) string {
	m.RLock()
	if icons != nil {
		icon := icons[name]
		m.RUnlock()
		return icon
	}
	m.RUnlock()

	// 获取互斥锁
	m.Lock()
	if icons == nil { // 必须重新检查nil值
		loadIcons()
	}
	icon := icons[name]
	m.Unlock()
	return icon
}

func main() {
	fmt.Println(icon("spades.png"))
	fmt.Println(icon("hearts.png"))
}
