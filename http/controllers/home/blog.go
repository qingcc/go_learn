package home

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ArticleList(c *gin.Context) {
	c.HTML(http.StatusOK, "/index", gin.H{})
}

func ArticleDetail(c *gin.Context) {

}
