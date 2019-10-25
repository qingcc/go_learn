package middleware

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/url"
	"os"
	"time"
	"github.com/qingcc/go_learn/util"
)

//region Remark: 将请求的日志记录并保存下来 Author; chijian
func Web(model string) gin.HandlerFunc {

	return func(c *gin.Context) {
		requesturl := c.Request.URL.String()
		method := c.Request.Method
		bodyBytes, _ := ioutil.ReadAll(c.Request.Body)
		body, _ := url.QueryUnescape(string(bodyBytes))
		connection := c.Request.Header.Get("Connection")
		upgrade := c.Request.Header.Get("Upgrade-Insecure-Requests")
		referer := c.Request.Header.Get("Referer")
		accept_encoding := c.Request.Header.Get("Accept-Encoding")
		accept_language := c.Request.Header.Get("Accept-Language")
		cache_control := c.Request.Header.Get("Cache-Control")
		user_agent := c.Request.Header.Get("User-Agent")
		accept := c.Request.Header.Get("Accept")
		cookie := c.Request.Header.Get("Cookie")
		//client_ip := http.ClientIp(c)	// 获取客户端IP 函数
		//public_ip := http.ClientIp(c)	// 获取客户端IP 函数

		var str = ""
		str = str + requesturl + ","
		str = str + method + ","
		str = str + body + ","
		//str = str + client_ip + ","
		//str = str + public_ip + ","
		str = str + connection + ","
		str = str + upgrade + ","
		str = str + referer + ","
		str = str + accept_encoding + ","
		str = str + accept_language + ","
		str = str + cache_control + ","
		str = str + user_agent + ","
		str = str + accept + ","
		str = str + cookie + ","

		//创建文件夹
		path := "./uploads/logs/" + model + "/" + time.Now().Format("2006/0102/")
		util.DirectoryMkdir(path)
		file := path + time.Now().Format("2006_01_02_15_00_00") + ".csv"
		request_logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
		if err != nil {
			fmt.Println(err.Error())
		}
		defer request_logFile.Close()
		request_logFile.WriteString(str)

		c.Request.Body.Close()
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		//fmt.Println(str)
	}
}

//endregion
