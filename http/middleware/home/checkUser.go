package middlewares

import (
	"github.com/qingcc/go_learn/config"
	"github.com/gin-gonic/gin"
	"github.com/qingcc/go_learn/logic"
	"net/http"
	"strconv"
	"github.com/qingcc/go_learn/util"
)

func CheckLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		//判断id对应的用户是否存在
		uid, _ := strconv.ParseInt(c.PostForm("keyid"), 10, 64)
		_uid := (uid + 1) / 3

		user := logic.UserLogic{}.FindOne(c, "id", _uid)
		if user.Id == 0 {
			c.JSON(http.StatusOK, gin.H{
				"status": config.HttpError,
				"msg":    "用户不存在",
			})
			return
		}
		token := c.PostForm("token")
		if token == "" {
			c.JSON(http.StatusOK, gin.H{
				"status": config.HttpError,
				"msg":    "登录超时",
			})
			return
		}
		_token, _ := util.AesEncrypt{}.AesDecrypt([]byte(token), []byte(user.Key[:8]+user.Token[8:]))
		if user.Token != string(_token) {
			c.JSON(http.StatusOK, gin.H{
				"status": config.HttpError,
				"msg":    "登录超时",
			})
			return
		}

		c.Request.Form.Set("keyid", strconv.FormatInt(_uid, 10))
		c.Next()
	}
}
