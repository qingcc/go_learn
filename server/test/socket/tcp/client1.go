package tcp

import (
	"fmt"
	"net"
)

func main() {
	for {
		readAndWriteClient() //成功连接上服务器
	}
}

const (
	proSplitChar = '|'
)

//成功连接上服务器
func readAndWriteClient() {
	conn, err := net.Dial("tcp", ":3366")
	if err != nil {
		//handle error
		fmt.Println(err.Error())
	}
	defer conn.Close()

	//time.Sleep(time.Second * 2)
	//data := "abcdefghij12345"
	data := make([]byte, 25660)
	n, err := conn.Read(data)

	fmt.Println(string(data[:n-1]))
}
