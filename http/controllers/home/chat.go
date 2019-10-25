package home

import (
	"github.com/qingcc/blog_go/config"
	"github.com/qingcc/blog_go/databases"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/qingcc/blog_go/logic"
	"github.com/qingcc/blog_go/model"
	"net/http"
	"strconv"
	"strings"
	"github.com/qingcc/blog_go/util"
)

func Chat(c *gin.Context) {
	user_id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	user := logic.UserLogic{}.FindOne(c, "id", user_id)
	imgHost := logic.ConfigLogic{}.ReadConfig("img_ip")

	item := make(map[string]interface{})
	item["username"] = user.Username
	item["img"] = imgHost + user.Url
	ws_ip := logic.ConfigLogic{}.ReadConfig("ws_ip")
	item["ws_ip"] = ws_ip + "/index/" + c.Query("id") + "/" + user.Token
	item["id"] = user_id
	item["token"] = user.Token
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.HTML(http.StatusOK, "/chat", gin.H{
		"data": "",
		"item": item,
	})
}
func GetData(c *gin.Context) {
	user_id, _ := strconv.ParseInt(c.PostForm("keyid"), 10, 64)

	user := logic.UserLogic{}.FindOne(c, "id", user_id)
	imgHost := logic.ConfigLogic{}.ReadConfig("img_ip")
	item := make(map[string]interface{})
	item["username"] = user.Username
	item["img"] = imgHost + user.Url
	ids := strings.Split(user.Url, "\\")
	_ids := ids[len(ids)-1:]
	item["_img"] = _ids[0]
	item["id"] = user_id
	item["token"] = user.Token
	data := logic.ChatLogic{}.GetData(10)
	list := make([]interface{}, len(data))
	for key, value := range data {
		u := logic.UserLogic{}.FindOne(c, "id", value.Uid)
		one := make(map[string]interface{})
		one["id"] = value.Cid
		one["content"] = value.Content
		one["url"] = imgHost + u.Url
		one["username"] = u.Username
		one["uid"] = value.Uid
		list[key] = one
	}

	_users := make([]interface{}, 0)
	redis_key := "ws_links"
	value, _ := redis.String(util.Get(redis_key))
	if value != "" {
		users := logic.UserLogic{}.FindIn(c, strings.Split(value, ","))
		if len(users) > 1 {
			for _, value := range users {
				item := make(map[string]interface{})
				item["id"] = value.Id
				item["username"] = value.Username
				item["url"] = imgHost + value.Url
				_users = append(_users, item)
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpSuccess,
		"data":   list,
		"user":   item,
		"users":  _users,
	})
	return
}

func PostChat(c *gin.Context) {
	user_id, _ := strconv.ParseInt(c.PostForm("keyid"), 10, 64)
	content := c.PostForm("content")
	if content == "" {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"msg":    "内容不能为空",
		})
		return
	}
	item := &model.Chat{Content: content, Uid: user_id}

	has, err := databases.Orm.Insert(item)
	util.CheckErr(err)

	if has < 1 {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"msg":    "",
		})
		return
	}
	msg := map[string]interface{}{
		"keyid": user_id,
		"token": c.PostForm("token"),
		"id":    item.Cid,
	}
	broadcast <- msg
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpSuccess,
		"msg":    "",
	})
	return
}

func WsChat(c *gin.Context) {
	user_id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	user := logic.UserLogic{}.FindOne(c, "id", user_id)
	imgHost := logic.ConfigLogic{}.ReadConfig("img_ip")

	item := make(map[string]interface{})
	item["username"] = user.Username
	item["img"] = imgHost + user.Url
	ws_ip := logic.ConfigLogic{}.ReadConfig("ws_ip")
	item["ws_ip"] = ws_ip + "/index/" + c.Query("id") + "/" + user.Token

	item["id"] = user_id
	item["token"] = user.Token
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.HTML(http.StatusOK, "/w_chat", gin.H{
		"data": "",
		"item": item,
	})
}
