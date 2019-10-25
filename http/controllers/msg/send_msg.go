package ws

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/qingcc/blog_go/logic"
	"net"
)

var (
	group     = make(map[string]map[string]*websocket.Conn)
	all       = make(map[string]*websocket.Conn)
	broadcast = make(chan map[string]interface{})
	//clientid  = make(chan *websocket.Conn)
	addUser = make(chan map[string]int, 20)
)

func init() {
	go Conn()
}

func Conn() {
	connAddr := logic.ConfigLogic{}.ReadConfig("tcp_ip")
	l, err := net.Listen("tcp", connAddr)
	if err != nil {
		fmt.Println("listen failed:", err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err.Error())
		}
		for {
			buf := make([]byte, 1024)
			n, err := conn.Read(buf)
			if err != nil {
				fmt.Println("tcp 连接 读取信息错误:", err)
			}
			msg := make(map[string]interface{})
			err = json.Unmarshal(buf[:n], msg)
			if all[msg["id"].(string)] != nil {

			}
		}
	}
}

func addOrDeleteUser(msg map[string]interface{}) map[string]interface{} {
	imgHost := logic.ConfigLogic{}.ReadConfig("img_ip")
	add, del := make([]interface{}, 0), make([]interface{}, 0)
	for key, value := range <-addUser {
		user := logic.UserLogic{}.GetOne("id", key)
		item := make(map[string]interface{})
		item["id"] = user.Id
		item["url"] = imgHost + user.Url
		item["username"] = user.Username
		if value == 1 { //添加
			add = append(add, item)
		} else if value == -1 {
			del = append(del, item)
		}
	}
	if len(add) != 0 {
		msg["add"] = add
	}
	if len(del) != 0 {
		msg["del"] = del
	}
	return msg
}
