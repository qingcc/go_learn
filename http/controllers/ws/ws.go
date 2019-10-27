package ws

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/qingcc/goblog/util"
	"net/http"
)

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Wshandler(c *gin.Context) {
	user_id := c.Param("keyid")

	ws, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("failed to set websocket upgrade ", err)
		return
	}
	client[user_id] = ws
	util.AddWsLinksToRedis(user_id)
	addUser <- map[string]int{user_id: 1}
}
