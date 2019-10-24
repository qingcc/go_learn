package tcp

import (
	"fmt"
	"net"
	"time"
)

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		// read from the connection
		// ... ...
		fmt.Println("accept connect")
		time.Sleep(time.Second * 100)

		// write to the connection
		//... ...
	}
}

func main() {
	//simple()

	//对方服务的listen backlog满
	backLogFull()
}

func simple() {
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			break
		}
		// start a new goroutine to handle
		// the new connection.
		go handleConn(c)
	}
}

func backLogFull() {
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("error listen:", err)
		return
	}
	defer l.Close()
	fmt.Println("listen ok")

	var i int
	for {
		time.Sleep(time.Second * 10)
		if _, err := l.Accept(); err != nil {
			fmt.Println("accept error:", err)
			break
		}
		i++
		fmt.Printf("%d: accept a new connection time: %s\n", i, time.Now().String())
	}
}
