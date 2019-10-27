package home

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qingcc/goblog/logic"
	"net/http"
)

func Index(c *gin.Context) {
	cate := logic.CategoryLogic{}.LevelList(c)
	for _, value := range cate {
		fmt.Println("cate:", value)
	}

	c.HTML(http.StatusOK, "/blog/home", gin.H{
		"Data":       logic.ArticleLogic{}.All(c),
		"Categories": cate,
	})
	return
}

func ArticleDetail(c *gin.Context) {

}
