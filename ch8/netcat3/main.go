package main

import (
	"io"
	"log"
	"net"
	"os"
)

// 客户端收到输入后立即退出,和服务器连接没啥关系
func main(){
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil{
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func(){
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{}  // 指示主goroutine
	}()
	mustCopy(conn, os.Stdin)
	//类型断言，调用*net.TCPConn的方法CloseWrite()只关闭TCP的写连接 ex8.3
	cw := conn.(*net.TCPConn)
	cw.CloseWrite()
	<-done // 等待后台goroutine完成
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
