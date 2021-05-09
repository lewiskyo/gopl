package main

import (
	"fmt"
	"sync"
)

var icons map[string]string
var loadIconsOnce sync.Once

func loadIcons() {
	icons = map[string]string{
		"spades.png":   "spades.png",
		"hearts.png":   "hearts.png",
		"diamonds.png": "diamonds.png",
		"clubs.png":    "clubs.png",
	}
}

func icon(name string) string {
	loadIconsOnce.Do(loadIcons)
	return icons[name]
}

func main() {
	fmt.Println(icon("spades.png"))
	fmt.Println(icon("hearts.png"))
}
