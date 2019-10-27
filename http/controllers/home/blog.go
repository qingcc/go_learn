package home

import (
	"github.com/gin-gonic/gin"
	"github.com/qingcc/goblog/logic"
	"net/http"
	"strconv"
)

func Index(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	category_id, _ := strconv.ParseInt(c.Query("category"), 10, 64)
	tag_id := c.Query("tag")
	time := c.Query("time")
	limit := 2
	c.HTML(http.StatusOK, "/blog/home", gin.H{
		"Data":       logic.ArticleLogic{}.All(c, page, limit, category_id, tag_id, time),
		"Categories": logic.GetAllCategory(0),
	})
	return
}

func ArticleDetail(c *gin.Context) {

}
