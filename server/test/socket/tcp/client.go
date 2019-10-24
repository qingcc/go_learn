package tcp

import (
	"fmt"
	"net"
	"time"
)

func main() {
	connectSuccessfully() //成功连接上服务器

	//连接的建立可能遇到的情况:
	// 1. 网络不可达或对方服务未启动 //服务器未监听该端口
	//conn_err1()	//dial tcp :222: connectex: No connection could be made because the target machine actively refused it.

	// 2. 对方服务的listen backlog满
	//conn_err2()
}

//成功连接上服务器
func connectSuccessfully() {
	conn, err := net.Dial("tcp", ":1200")
	if err != nil {
		//handle error
		fmt.Println(err.Error())
	}

	conn.Write([]byte("hello"))
	fmt.Println(conn)
}

// 1. 网络不可达或对方服务未启动 //服务器未监听该端口
func conn_err1() {
	conn, err := net.Dial("tcp", ":222")
	if err != nil {
		//handle error
		fmt.Println(err.Error()) //dial tcp :222: connectex: No connection could be made because the target machine actively refused it.
	}
	fmt.Println(conn)
}

func conn_err2() {
	// 2. 对方服务的listen backlog满

	var sl []net.Conn
	for i := 1; i < 1000; i++ {
		conn := establishConn(i)
		if conn != nil {
			sl = append(sl, conn)
		}
	}

	time.Sleep(time.Second * 10000)
}

func establishConn(i int) net.Conn {

	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		//handle error
		fmt.Println(err.Error(), " ", i)
	}
	return conn
}

// 3、网络延迟较大，Dial阻塞并超时, 未测试,需模拟一个延迟较大的网络环境
func conn_err3() {
	conn, err := net.Dial("tcp", ":222")
	if err != nil {
		//handle error
		fmt.Println(err.Error()) //dial tcp :222: connectex: No connection could be made because the target machine actively refused it.
	}
	fmt.Println(conn)
}
