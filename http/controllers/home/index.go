package home

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qingcc/goblog/config"
	"github.com/qingcc/goblog/databases"
	"github.com/qingcc/goblog/logic"
	"github.com/qingcc/goblog/model"
	"github.com/qingcc/goblog/util"
	"net/http"
	"time"
)

func Reg(c *gin.Context) {
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func GetLogin(c *gin.Context) {
	//test()
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.HTML(http.StatusOK, "/login", gin.H{})
}

func PutLogin(c *gin.Context) {
	name := c.PostForm("name")
	pass := c.PostForm("pass")
	println("name:", name, "pass:", pass)
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpError,
		"msg":    "put提交",
	})
	return

}

func DeleteLogin(c *gin.Context) {
	mul := c.Param("name")
	ch := make(chan int, 10)
	fmt.Println(ch)
	name, _ := c.Get("name")
	pass := c.PostForm("pass")
	fmt.Println("delete name:", name, "pass:", pass, "mul:", mul)
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpError,
		"msg":    "deletet提交",
	})
	return

}

func test() {
	user := new(model.Test1)
	user.Title = "test1"
	user.Content = "test@test.5416341fsdfs563df1sfcom"
	user.High = 3.258
	user.CoinId = 111111
	user.Tim = time.Now().AddDate(0, 1, 1)
	user.IsSys = true
	has, err := databases.Orm.Id(1).Update(user)
	fmt.Println(has, err, user)
}

func PostLogin(c *gin.Context) {
	name := c.PostForm("name")
	pass := c.PostForm("pass")
	user := logic.UserLogic{}.FindOneUser(c, name)

	password, _ := util.AesEncrypt{}.AesEncrypt([]byte(pass), []byte(user.Key))
	if string(password) != user.Password {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"msg":    "用户名或密码错误!",
		})
		return
	}
	token := util.GetSjCode(16)
	user.Token = token
	has, _ := databases.Orm.Id(user.Id).Cols("token").Update(user)
	if has < 1 {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"msg":    "登录失败",
		})
		return
	}
	item := user.EncryptToken()

	go func() {
		logic.ToConn("tcp", item)
	}()

	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpSuccess,
		"msg":    "登录成功",
		"data":   item,
	})
	return
}

func PostReg(c *gin.Context) {
	name := c.PostForm("name")
	pass := c.PostForm("pass")
	email := c.PostForm("email")
	user := logic.UserLogic{}.FindOne(c, "name", name)
	if user.Id != 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"msg":    "用户名已注册",
		})
		return
	}
	if name == "" {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"msg":    "用户名不能为空",
		})
		return
	}
	if pass == "" {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"msg":    "密码不能为空",
		})
		return
	}
	user = logic.UserLogic{}.FindOne(c, "email", email)
	if user.Id != 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"msg":    "邮箱已注册",
		})
		return
	}

	key := util.GetSjCode(16)
	password, _ := util.AesEncrypt{}.AesEncrypt([]byte(pass), []byte(key))
	url := logic.ConfigLogic{}.ReadConfig("default_url")

	has, err := databases.Orm.Insert(model.User{Username: name, Email: email, Password: string(password), Key: key, Url: url})
	if err != nil {
		fmt.Println(err.Error())
	}
	if has < 1 {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"msg":    "注册失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpSuccess,
		"msg":    "注册成功",
	})
	return
}
