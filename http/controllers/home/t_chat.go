package home

import (
	"github.com/gin-gonic/gin"
	"github.com/qingcc/goblog/config"
	"github.com/qingcc/goblog/databases"
	"github.com/qingcc/goblog/logic"
	"github.com/qingcc/goblog/model"
	"github.com/qingcc/goblog/util"
	"net/http"
	"strconv"
)

func GetTData(c *gin.Context) {
	user_id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	user := logic.UserLogic{}.FindOne(c, "id", user_id)
	imgHost := logic.ConfigLogic{}.ReadConfig("img_ip")
	item := make(map[string]interface{})
	item["username"] = user.Username
	item["img"] = imgHost + user.Url
	item["id"] = user_id
	item["token"] = user.Token
	data := logic.TChatLogic{}.GetData(10)
	list := make([]interface{}, len(data))
	for key, value := range data {
		u := logic.UserLogic{}.FindOne(c, "id", value.Id)
		one := make(map[string]interface{})
		one["id"] = value.Id
		one["content"] = value.Content
		one["url"] = imgHost + u.Url
		one["username"] = u.Username
		one["uid"] = value.Uid
		one["to_uid"] = value.ToUid
		list[key] = one
	}
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpSuccess,
		"data":   list,
		"user":   item,
	})
	return
}

func PostTChat(c *gin.Context) {
	user_id, _ := strconv.ParseInt(c.Param("keyid"), 10, 64)
	content := c.PostForm("content")
	to_uid, _ := strconv.ParseInt(c.PostForm("to_uid"), 10, 64)
	if content == "" {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"msg":    "内容不能为空",
		})
		return
	}
	item := &model.TChat{Uid: user_id, Content: content, ToUid: to_uid}
	has, err := databases.Orm.Insert(item)
	util.CheckErr(err)

	if has < 1 {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"msg":    "",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpSuccess,
		"msg":    "",
	})
	return
}

func TChat(c *gin.Context) {
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
	c.HTML(http.StatusOK, "/t_chat", gin.H{
		"data": "",
		"item": item,
	})
}
