package controllers

import (
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
)

var captchaHandler = captcha.Server(100, 40)

//验证码

func Catpcha(c *gin.Context) {
	Server(c)
}

func Server(c *gin.Context) error {
	captchaHandler.ServeHTTP(c.Writer, c.Request)
	return nil
}
