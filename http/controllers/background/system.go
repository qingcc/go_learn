package background

import (
	"github.com/qingcc/go_learn/config"
	"github.com/gin-gonic/gin"
	"github.com/qingcc/go_learn/logic"
	"github.com/qingcc/go_learn/model"
	"net/http"
)

//region Remark: 系统基础设置 Author:Qing
func GetSystemBase(c *gin.Context) {

	c.HTML(http.StatusOK, "system/base", gin.H{
		"title":          "Background Login",
		"sys_name":       logic.ConfigLogic{}.ReadConfig("sys.name"),
		"sys_is_open":    logic.ConfigLogic{}.ReadConfig("sys.is_open"),
		"sys_close_info": logic.ConfigLogic{}.ReadConfig("sys.close_info"),
		"sys_cache_time": logic.ConfigLogic{}.ReadConfig("sys.cache_time"),
		"ws_ip":          logic.ConfigLogic{}.ReadConfig("ws_ip"),
	})
}
func PostSystemBase(c *gin.Context) {
	data := []model.Config{
		model.Config{
			Name:  "ws_ip",
			Value: c.DefaultPostForm("ws_ip", ""),
		},
	}
	ws_ip := c.PostForm("ws_ip")

	if ws_ip == "" {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"msg":    "推送IP不能为空",
		})
		return
	}
	logic.ConfigLogic{}.SetConfig(data)
	//c.JSON(http.StatusOK, gin.H{
	//	"status": config.HttpSuccess,
	//	"info":   "保存成功",
	//	"url":    "/admin/config",
	//})

	c.HTML(http.StatusOK, "/system/base", gin.H{
		"status": config.HttpSuccess,
		"info":   "保存成功",
		"url":    "/admin/system/base",
	})
}

//endregion

//region Remark: 系统屏蔽词 Author:Qing
func GetSystemShield(c *gin.Context) {
	c.HTML(http.StatusOK, "/system/shield", gin.H{
		"title": "BackGround Center",
		"info":  "",
	})
}

//endregion

//region Remark: 系统日志 Author:Qing
func GetSystemLog(c *gin.Context) {
	c.HTML(http.StatusOK, "/system/log", gin.H{
		"title": "BackGround Center",
		"info":  "",
	})
}

//endregion
