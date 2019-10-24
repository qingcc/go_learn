package logic

import (
	"encoding/json"
	"fmt"
	"net"
)

func ToConn(network string, item map[string]interface{}) {
	connAddr := ConfigLogic{}.ReadConfig("tcp_ip")
	conn, err := net.Dial(network, connAddr)
	if err != nil {
		fmt.Println("tcp connect to server failed:", err)
	}

	go func() { //连接后写入认证信息
		if item["uid"].(int64) != 0 && item["token"].(string) != "" {
			_item, err := json.Marshal(item)
			if err != nil {
				fmt.Println(err.Error())
			}
			_, err = conn.Write([]byte(_item))
			if err != nil {
				fmt.Println("写入失败")
			}
		}
	}()
}
