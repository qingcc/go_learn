package middleware

import (
	"github.com/qingcc/blog_go/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qingcc/blog_go/logic"
	"net/http"
	"github.com/qingcc/blog_go/util"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//region Remark: 判断session里是否存在 adminid Author:Qing
		if util.HasSession(c, "adminid") == false {
			fmt.Println("未登录,请先登录")
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
		}
		//endregion

		//region Remark: 判断权限 Author:Qing
		nowLoginManagerInfo := logic.DefaultAdmin.AdminNowLoginInfo(c)
		hander_name := c.HandlerName()
		//判断当前用户是否有权限访问该资源---如果不是超级管理员，则进行权限节点的校验
		if nowLoginManagerInfo.Role.IsSuper == false {
			if logic.DefaultRoleNodeRoutes.HasAuth(nowLoginManagerInfo.RoleNodesRoutes, hander_name) == false {
				//失败
				fmt.Println(222)
				if c.Request.Method == "GET" {
					c.JSON(http.StatusOK, gin.H{
						"status": config.HttpError,
						"info":   "您没有管理该页面的权限，请勿非法进入！",
					})
				} else {
					c.JSON(http.StatusOK, gin.H{
						"status": config.HttpError,
						"info":   "您没有操作该页面的权限，请勿非法进入！",
					})
				}
				c.Abort()
			}
		}

		//if res := logic.DefaultAdminNavigationNode.FindOne(c, "route_action", hander_name); res == nil {
		//	c.JSON(http.StatusOK, gin.H{
		//		"status": config.HttpError,
		//		"info":   hander_name + "没有在数据库中存在，请添加后再操作",
		//	})
		//	c.Abort()
		//}
		//endregion

		//谷歌验证码

		//写入日志

		c.Next()
	}
}
