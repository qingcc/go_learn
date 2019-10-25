package routers

import (
	"github.com/qingcc/go_learn/databases"
	"github.com/gin-gonic/gin"
	"github.com/qingcc/go_learn/http/controllers/ws"
)

func InitWsRouter() *gin.Engine {
	router := gin.Default()

	databases.Init()
	router.GET("/index/:keyid/:token", ws.Wshandler)
	//r := router.Group("/user", middlewares.CheckLogin())
	router.POST("/user/send", ws.Send)

	return router
}
