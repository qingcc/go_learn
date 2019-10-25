package routers

import (
	"github.com/qingcc/goblog/databases"
	"github.com/foolin/gin-template"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/qingcc/goblog/http/controllers/home"
	"github.com/qingcc/goblog/http/middleware/home"
)

func InitHomeRouter() *gin.Engine {
	router := gin.Default()
	router.Use(sessions.Sessions("go_home", databases.SessionStore()))

	databases.Init()
	//region Remark: 定义公共的中间件 Author; chijian
	//endregion
	router.Static("/public", "./static/home") //上传路径
	router.Static("/uploads", "./uploads")    //上传路径
	router.HTMLRender = gintemplate.New(gintemplate.TemplateConfig{
		Root:      "resources/home",
		Extension: ".html",
		//Master:       "layouts/master",
		DisableCache: true,
	})
	//router.Use(middleware.Web("background"))
	router.GET("/login", home.GetLogin)
	router.PUT("/login", home.PutLogin)
	router.DELETE("/login", home.DeleteLogin)
	router.POST("/login", home.PostLogin)
	router.GET("/reg", home.Reg)
	router.POST("/reg", home.PostReg)
	//router.GET("/captcha", controllers.Catpcha)
	router.GET("/user/chat", home.Chat)
	router.GET("/user/t_chat", home.WsChat) //websocket

	router.GET("/blog_list", home.ArticleList)

	user := router.Group("/user", middlewares.CheckLogin())
	//user.GET("/center", background.GetCenter)

	user.POST("/to_chat", home.PostChat)
	user.POST("/get_data", home.GetData)

	return router
}
