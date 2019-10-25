package main

import (
	"github.com/gin-gonic/gin"
	"github.com/qingcc/go_learn/routers"
)

func main() {
	gin.SetMode(gin.DebugMode)
	router := routers.InitBackGroundRouter()
	//router := gin.Default()
	//defer databases.Orm.Close()
	//_ = databases.Orm.Sync2(new(model.AdminNavigation), new(model.AdminNavigationNode), new(model.Admin), new(model.Role), new(model.RoleNode), new(model.Category), new(model.Article), new(model.RoleNodeRoutes))
	//databases.Orm.Sync2(new(model.User), new(model.Config))
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//logger.Init(config.ROOT+"/log", config.ConfigFile.MustValue("global", "log_level", "DEBUG"))
	//println("啊啊啊")
	router.Run(":8049")

	//go ServerBackGround()
}

//func ServerBackGround() {
//
//	//常驻内存的数据
//	//go loadData()
//}
//
//func loadData() {
//	logic.LoadRoleAuthorities()
//}
