package routers

import (
	"blog_go/databases"
	"github.com/foolin/gin-template"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	//_ "github.com/go-sql-driver/mysql"
	"blog_go/http/controllers/background"
	"blog_go/http/middleware"
)

func InitBackGroundRouter() *gin.Engine {
	router := gin.Default()
	router.Use(sessions.Sessions("go", databases.SessionStore()))

	databases.Init()
	//region Remark: 定义公共的中间件 Author; chijian
	//endregion
	router.Static("/public", "./static")   //渲染html页面//router.LoadHTMLGlob("resources/**/*")
	router.Static("/uploads", "./uploads") //上传路径
	router.HTMLRender = gintemplate.New(gintemplate.TemplateConfig{
		Root:         "resources/background",
		Extension:    ".html",
		Master:       "layouts/master",
		DisableCache: true,
	})
	router.Use(middleware.Web("background"))
	router.GET("/login", background.GetLogin)
	router.POST("/login", background.PostLogin)
	//router.GET("/captcha", controllers.Catpcha)

	//测试
	router.POST("/test", background.PostTest)

	v1 := router.Group("/admin", middleware.Auth())
	v1.GET("/center", background.GetCenter)
	v1.GET("/welcome", background.GetWelcome)
	v1.GET("/clear", background.Clear)
	v1.GET("/logout", background.Logout)
	v1.GET("/icon/list", background.Icon)
	//v1.POST("/upload/controllers/image", util.WangEditorUpload)
	//v1.POST("/upload/image", util.UploadImage)
	//v1.POST("/upload/file", util.UploadFile)
	//v1.POST("/upload/video", util.UploadVideo)

	//数据备份
	v1.GET("/backup/list", background.BackUpList)
	v1.GET("/backup/do", background.BackUp)
	v1.GET("/backup/download/:name", background.DownloadBackup)
	v1.GET("/backup/del/:name", background.DelBackup)

	//管理员
	v1.GET("/admin_user/list", background.GetAdminUserList)
	v1.GET("/admin_user/add", background.GetAdminUserAdd)
	v1.POST("/admin_user/add", background.PostAdminUserAdd)
	v1.GET("/admin_user/edit", background.GetAdminUserEdit)
	v1.POST("/admin_user/edit", background.PostAdminUserEdit)
	v1.POST("/admin_user/del", background.PostAdminUserDel)
	v1.POST("/admin_user/enable", background.PostAdminUserEnable)

	//角色
	v1.GET("/admin_role/list", background.GetRoleList)
	v1.GET("/admin_role/add", background.GetRoleAdd)
	v1.POST("/admin_role/add", background.PostRoleAdd)
	v1.GET("/admin_role/edit", background.GetRoleEdit)
	v1.POST("/admin_role/edit", background.PostRoleEdit)
	v1.POST("/admin_role/del", background.PostRoleDel)

	//权限节点
	v1.GET("/navigation/list", background.GetNavigationList)
	v1.GET("/navigation/add", background.GetNavigationAdd)
	v1.POST("/navigation/add", background.PostNavigationAdd)
	v1.GET("/navigation/edit", background.GetNavigationEdit)
	v1.POST("/navigation/edit", background.PostNavigationEdit)
	v1.POST("/navigation/sort", background.PostNavigationSort)
	v1.POST("/navigation/del", background.PostNavigationDel)

	//系统
	v1.GET("/system/base", background.GetSystemBase)
	v1.POST("/system/base", background.PostSystemBase)
	v1.GET("/system/shield", background.GetSystemShield)
	v1.GET("/system/log", background.GetSystemLog)

	//博文分类
	v1.GET("/category/list", background.GetCateList)
	v1.GET("/category/add", background.GetCateAdd)
	v1.POST("/category/add", background.PostCateAdd)
	v1.GET("/category/edit", background.GetCateEdit)
	v1.POST("/category/edit", background.PostCateEdit)
	v1.POST("/category/del", background.PostCateDel)
	v1.POST("/category/sort", background.PostCateSort)

	//博文管理
	v1.GET("/article/list", background.GetArticleList)
	v1.GET("/article/add", background.GetArticleAdd)
	v1.POST("/article/add", background.PostArticleAdd)
	v1.GET("/article/edit", background.GetArticleEdit)
	v1.POST("/article/edit", background.PostArticleEdit)
	//v1.POST("/article/del", background.PostArticleDel)
	v1.POST("/article/sort", background.PostArticleSort)
	v1.POST("/article/del", new(background.ArticleController).Del)

	return router
}
