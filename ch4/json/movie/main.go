package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// 只有首字母大写的成员才可以转换为JSON字段
type Movie struct {
	Title  string
	Year   int  `json:"released"`        // 成员标签
	Color  bool `json:"color,omitempty"` // 成员标签, omitempty表示这个成员的值是零值或者为空,则不输出到JSON中
	Actors []string
}

var movies = []Movie{
	{Title: "Ma", Year: 1942, Color: false, Actors: []string{"aa", "bb"}},
	{Title: "Mb", Year: 1988, Color: true, Actors: []string{"aa"}},
	{Title: "Mc", Year: 1941, Color: false, Actors: []string{"aa", "bb"}},
}

func main() {
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("json marshal failed: %s", err)
	}
	fmt.Printf("%s\n\n", data)

	// 格式化输出,第二个参数表示每行输出的前缀字符串, 第三个参数定义缩进的字符串
	data2, err := json.MarshalIndent(movies, "", "   ")
	if err != nil {
		log.Fatalf("json marshalindent failed: %s", err)
	}
	fmt.Printf("%s\n\n", data2)

	// 结构体唯一成员是Title,只把JSON数据中Title字段解码,其他数据丢弃
	var titles []struct{ Title string }
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON unmarshaling failed %s", err)
	}
	fmt.Println(titles) // [{Ma} {Mb} {Mc}]
}
