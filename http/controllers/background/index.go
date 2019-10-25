package background

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"github.com/qingcc/goblog/util"
)

//region Remark: 主页 Author:Qing
func GetCenter(c *gin.Context) {
	c.HTML(http.StatusOK, "/index", gin.H{
		"title": "BackGround Center",
		"info":  "",
	})
}

//endregion

//region Remark: 主页 Author:Qing
func GetWelcome(c *gin.Context) {
	//var list map[string]interface{}
	//list["goos"] = runtime.GOOS         //系统版本
	//list["hostName"], _ = os.Hostname() //主机名
	//list["goarch"] = runtime.GOARCH     //系统构架 如amd64
	//list["numCPU"] = runtime.NumCPU()   //cpu数量

	c.HTML(http.StatusOK, "layouts/welcome", gin.H{
		"Title": "BackGround Center",
		"info":  "",
	})
}

//endregion

//region Remark: 登出 Author:Qing
func Clear(c *gin.Context) {
	fmt.Println("clear")
	util.DelAll()
	c.HTML(http.StatusOK, "/login", gin.H{
		"Title": "BackGround Login",
		"info":  "清空缓存成功",
		"url":   "/admin/center",
	})
}

//endregion

//region Remark: 登出 Author:Qing
func Logout(c *gin.Context) {
	util.DelSession(c, "adminid")
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"info":   "已成功登出",
		"url":    "/login",
	})
}

//endregion

//region Remark: 图标 Author:Qing
func Icon(c *gin.Context) {
	//s := make([]string, 0)
	//for i := 1546; i <= 1837; i++ {
	//	n := fmt.Sprintf("%x", i)
	//	str := "&#xe" + n + ";"
	//	s = append(s, str)
	//}
	//item := make(map[int]string, 7)
	//for key, value := range s {
	//	item[key%7] = value
	//	if key%7 == 0 && key != 0{
	//		WriteIcon2File(item)
	//	}
	//}
	c.HTML(http.StatusOK, "/icon", gin.H{
		"Title": "BackGround Login",
	})
}

//endregion

//region Remark: 将图标写入文件 Author:Qing
func WriteIcon2File(data map[int]string) {
	file, _ := os.OpenFile("./static/icon.html", os.O_WRONLY|os.O_APPEND, 0666)
	str := "<tr>"
	for _, value := range data {
		str = str + "<td>" +
			"<a href='javascript:;' class='ml-5' style='text-decoration:none'>" +
			"<i class='Hui-iconfont'> " + value + "</i>" +
			"</a>" +
			"<i> " + value[1:] + "</i>" +
			"</td>"
	}
	str = str + "</tr>\n"
	defer file.Close()
	n, err := file.WriteString(str)
	fmt.Println(n, err, data)
}

//endregion
