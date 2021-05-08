package main

import "fmt"

type ByteCounter int

// 仅统计传入数据的字节数
func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += (ByteCounter(len(p)))
	return len(p), nil
}

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c)

	c = 0
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)
}
