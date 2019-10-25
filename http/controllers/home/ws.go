/*
@Time : 2018/7/10 17:44
@Author : Administrator
@File : SendTrad
@Software: GoLand
*/
package home

import (
	"github.com/qingcc/blog_go/logic"
	"net/http"
	"net/url"
	"strconv"
)

var (
	broadcast = make(chan map[string]interface{})
)

func init() {
	go SendTrad()
}

func SendTrad() {
	for {
		a := <-broadcast

		user_id := a["keyid"].(int64)
		_user_id := strconv.FormatInt(user_id, 10)
		id := a["id"].(int64)
		_id := strconv.FormatInt(id, 10)

		//item := logic.ChatLogic{}.FindOneChat("id", id)
		//user := logic.UserLogic{}.GetOne("id", item.Uid)
		//imgHost := logic.ConfigLogic{}.ReadConfig("img_ip")

		ws_ip := logic.ConfigLogic{}.ReadConfig("ws_ip")

		_url := "http:" + ws_ip[3:] + "/user/send"

		http.PostForm(_url,
			url.Values{
				"keyid": {_user_id},
				"token": {a["token"].(string)},
				"id":    {_id},
			})
	}
}
