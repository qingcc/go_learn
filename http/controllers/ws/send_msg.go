package ws

import (
	"fmt"
	"github.com/gorilla/websocket"
	"logic"
)

var (
	//clients   = make(map[string]map[string]*websocket.Conn)
	client    = make(map[string]*websocket.Conn)
	broadcast = make(chan map[string]interface{})
	//clientid  = make(chan *websocket.Conn)
	addUser = make(chan map[string]int, 20)
)

func init() {
	go SendWs()
}

func SendWs() {
	for {
		msg := <-broadcast
		if len(client) == 0 { //无连接
			continue
		}
		//fmt.Println("len:", len(client), "client:", client)
		//for _, client := range clients {
		for u_id, ws := range client {
			msg = addOrDeleteUser(msg)
			err := ws.WriteJSON(msg) //推送消息到客户端
			if err != nil {
				fmt.Errorf("ws.WriteJSON err: %v", err)
				ws.Close()
				addUser <- map[string]int{u_id: -1}
				delete(client, u_id)
			}
		}
		//}
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
