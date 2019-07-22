package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()

	// 循环执行
	for {
		// 等待客户端连接, 阻塞
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		// 连接成功, 处理conn
		go handleConn(conn)
	}

}

// 只写的chan 类型是string
type client chan<- string

var (
	entering = make(chan client)
	//
	leaving = make(chan client)
	// 处理客户端发送来的消息
	message = make(chan string)
)

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-message: //如果message中有消息, 且成功读出到msg, 就将消息写入cli
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}

	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string)

	go writeToClient(conn, ch)

	who := conn.RemoteAddr().String()

	ch <- "you are " + who

	message <- who + "are arrived"

	entering <- ch
	// 接收一个输入
	input := bufio.NewScanner(conn)
	for input.Scan() {
		message <- who + ": " + input.Text()
	}

	leaving <- ch
	message <- who + " are left"
	conn.Close()
}

// 将消息发送给客户端
func writeToClient(conn net.Conn, ch <-chan string) {
	// 循环
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
