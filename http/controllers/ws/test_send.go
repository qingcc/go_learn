package ws

import (
	"github.com/gin-gonic/gin"
	"github.com/qingcc/blog_go/logic"
)

func Send(c *gin.Context) {
	id := c.PostForm("id")

	item := logic.ChatLogic{}.FindOne(c, "cid", id)

	user := logic.UserLogic{}.FindOne(c, "id", item.Uid)
	imgHost := logic.ConfigLogic{}.ReadConfig("img_ip")
	msg := map[string]interface{}{
		"uid":      item.Uid,
		"username": user.Username,
		"url":      imgHost + user.Url,
		"content":  item.Content,
	}
	broadcast <- msg

}
