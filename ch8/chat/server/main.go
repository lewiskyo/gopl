package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

type client chan<- string // 对外发送消息的通道

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // 所有接收的客户端消息
)

func broadcaster() {
	clients := make(map[client]bool) // 所有连接的客户端
	for {
		select {
		// 把所有接收的消息广播给所有的客户
		// 发送消息通道
		case msg := <-messages:
			for cli := range clients {
				cli <- msg
			}
			// 有客户连接
		case cli := <-entering:
			clients[cli] = true
			// 有客户端断开连接
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

const timeout = 20 * time.Second

func handleConn(conn net.Conn) {
	ch := make(chan string) // 对外发送客户端消息的通道
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "you are " + who
	messages <- who + " has arrived"
	entering <- ch

	// ex8.14 close后 input.scan会跳出循环
	timer := time.NewTimer(timeout)
	go func() {
		<-timer.C
		conn.Close()
	}()

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
		timer.Reset(timeout)
	}
	leaving <- ch
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
