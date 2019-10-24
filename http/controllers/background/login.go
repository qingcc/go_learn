package background

import (
	"databases"
	"github.com/gin-gonic/gin"
	"logic"
	"model"
	"net/http"
	"util"
)

func GetLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"Title": "BackGround Login",
		"info":  "info",
	})
}

func PostLogin(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	admin := logic.DefaultAdmin.GetAdminByField(c, "email", email)

	if admin.Id < 1 {
		c.JSON(http.StatusOK, gin.H{
			"status": 0,
			"info":   "无该用户",
			"aa":     admin,
		})
		return
	}
	if admin.IsLock == true {
		c.JSON(http.StatusOK, gin.H{
			"status": 0,
			"info":   "该账号已被锁定",
		})
		return
	}
	if admin.Password != util.Strmd5(password) {
		c.JSON(http.StatusOK, gin.H{
			"status": 0,
			"info":   "账号或密码不正确",
		})
		return
	}

	databases.Orm.Id(admin.Id).Incr("login_count").Update(&model.Admin{})

	//写入session 页面跳转
	util.SetSession(c, "adminid", admin.Id)

	c.JSON(http.StatusOK, gin.H{
		"status": 1,
		"info":   "登录成功",
		"url":    "/admin/center",
	})
	return
}
