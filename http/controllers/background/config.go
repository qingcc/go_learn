package background

import (
	"github.com/gin-gonic/gin"

	"github.com/qingcc/go_learn/config"
	"github.com/qingcc/go_learn/logic"
	"github.com/qingcc/go_learn/model"
	"net/http"
)

//region Remark: 列表 Author:Qing
func ReadConfig(c *gin.Context) {
	c.HTML(http.StatusOK, "config/config", gin.H{
		"title":          "Background Login",
		"sys_name":       logic.ConfigLogic{}.ReadConfig("sys.name"),
		"sys_is_open":    logic.ConfigLogic{}.ReadConfig("sys.is_open"),
		"sys_close_info": logic.ConfigLogic{}.ReadConfig("sys.close_info"),
		"sys_cache_time": logic.ConfigLogic{}.ReadConfig("sys.cache_time"),
		"ws_ip":          logic.ConfigLogic{}.ReadConfig("ws_ip"),
		"img_ip":         logic.ConfigLogic{}.ReadConfig("img_ip"),
	})
}

//endregion

//region Remark: 添加 Author:Qing
func PostConfig(c *gin.Context) {
	data := []model.Config{
		model.Config{
			Name:  "ws_ip",
			Value: c.DefaultPostForm("ws_ip", ""),
		},
		model.Config{
			Name:  "img_ip",
			Value: c.DefaultPostForm("img_ip", ""),
		},
		model.Config{
			Name:  "default_img",
			Value: c.DefaultPostForm("default_img", ""),
		},
	}
	ws_ip := c.PostForm("ws_ip")
	img_ip := c.PostForm("img_ip")
	default_img := c.PostForm("default_img")

	if ws_ip == "" {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"msg":    "推送IP不能为空",
		})
		return
	}
	if img_ip == "" {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"msg":    "图片地址IP不能为空",
		})
		return
	}
	if default_img == "" {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"msg":    "用户默认头像不能为空",
		})
		return
	}
	logic.ConfigLogic{}.SetConfig(data)
	//c.JSON(http.StatusOK, gin.H{
	//	"status": config.HttpSuccess,
	//	"info":   "保存成功",
	//	"url":    "/admin/config",
	//})

	c.HTML(http.StatusOK, "/admin/config", gin.H{
		"status": config.HttpSuccess,
		"info":   "保存成功",
		"url":    "/admin/config",
	})
}
