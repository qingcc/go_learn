package main

import (
	"github.com/gin-gonic/gin"
	"blog_go/routers"
)

func main() {
	gin.SetMode(gin.DebugMode)
	router := routers.InitHomeRouter()
	//err := databases.Orm.Sync2(new(model.Test1))
	//fmt.Println("err:", err)
	//defer databases.Orm.Close()
	router.Run(":2019")
}
