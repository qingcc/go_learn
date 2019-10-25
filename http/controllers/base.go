package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/polaris1119/logger"
	"github.com/qingcc/blog_go/logic"
)

//region Remark: 获取日志实例 Author:Qing
func getLogger(c *gin.Context) *logger.Logger {
	return logic.GetLogger(c)
}

//endregion

//region Remark: 排序,是否显示,删除 接口 Author:Qing
type BaseAdmin interface {
	PostSort()
	PostShow()
	PostDel()
}

//endregion
