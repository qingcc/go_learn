package tcp

import (
	"bufio"
	"fmt"
	"github.com/xiaogan18/msgserver"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	//simple()

	//对方服务的listen backlog满
	//ReadAndWrite()
	go server()
	test()
}

var msg = make(chan interface{})

func test() {
	for {
		// 从标准输入读取字符串，以\n为分割
		fmt.Println("input a msg:")
		text, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err == nil {
			text = strings.Replace(text, "\r\n", "", 1)
		}
		msg <- text
	}
}

func server() {
	sdr, lster, err := msgserver.NewDefaultServer(false) //不开启ssl加密传输
	if err != nil {
		fmt.Println(err)
	}
	//开启一个协程 监听主机3366端口
	go func() {
		lster.Listen("127.0.0.1:3366")
	}()
	sdr.BeginSender()

	for {
		//// 从标准输入读取字符串，以\n为分割
		//fmt.Println("input a msg:")
		//text, err := bufio.NewReader(os.Stdin).ReadString('\n')
		//if err == nil {
		//	text = strings.Replace(text, "\r\n", "", 1)
		//
		//	if text == "count" {
		//		fmt.Println(lster.OnlineCount())
		//	} else {
		//		sdr.SendNotice(text)
		//	}
		//}
		text := <-msg
		fmt.Println("text:", text)
		sdr.SendNotice(text)
	}
}

func ReadAndWrite() {
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
		fmt.Println("accept ok")
		// start a new goroutine to handle
		// the new connection.
		go readAndWrite(c)
	}
}

func readAndWrite(c net.Conn) {
	defer c.Close()
	for {
		var buf = make([]byte, 10)
		n, err := c.Read(buf)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println("read %d bytes, content is %s\n", n, string(buf[:n]))
		time.Sleep(time.Second * 2)
	}
}
