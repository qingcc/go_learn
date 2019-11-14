package background

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qingcc/goblog/config"
	"github.com/qingcc/goblog/databases"
	redisPool "github.com/qingcc/goblog/utils/redis"
	"time"

	//"github.com/polaris1119/logger"
	"github.com/polaris1119/logger"
	"github.com/qingcc/goblog/logic"
	"github.com/qingcc/goblog/model"
	"html/template"
	"net/http"
	"strconv"
)

//region Remark: 列表 Author:Qing
func GetArticleList(c *gin.Context) {
	c.HTML(http.StatusOK, "/article/list", gin.H{
		"Title": "BackGround Center",
		"Data":  logic.DefaultArticle.List(c),
	})
}

//endregion

//region Remark: 添加 Author:Qing
func GetArticleAdd(c *gin.Context) {
	c.HTML(http.StatusOK, "/article/add", gin.H{
		"Title": "BackGround Center",
		"Level": logic.DefaultCategory.LevelList(c),
	})
}

func PostArticleAdd(c *gin.Context) {
	title := c.PostForm("title")
	abstract := c.PostForm("abstract")
	//author := c.PostForm("author")
	cover := c.PostForm("cover")
	tags := c.PostForm("tags")
	sources := c.PostForm("sources")
	content := c.PostForm("content")
	sort, _ := strconv.ParseInt(c.PostForm("sort"), 10, 64)
	category_id, _ := strconv.ParseInt(c.PostForm("category_id"), 10, 64)
	is_show := true
	if c.PostForm("is_show") != "1" {
		is_show = false
	}
	allow_comments := true
	if c.PostForm("allow_comments") != "1" {
		allow_comments = false
	}
	if title == "" {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"msg":    "标题不能为空",
		})
		return
	}
	if cover == "" {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"msg":    "封面图不能为空",
		})
		return
	}
	if content == "" {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"msg":    "内容不能为空",
		})
		return
	}

	has, err := databases.Orm.Insert(&model.Article{
		Title:      title,
		Abstract:   abstract,
		CategoryId: category_id,
		//Author:        author,
		Cover:         cover,
		Tags:          tags,
		Sources:       sources,
		IsShow:        is_show,
		Sort:          sort,
		Year:          int64(time.Now().Year()),
		Month:         int64(time.Now().Month()),
		AllowComments: allow_comments,
		Content:       template.HTML(content),
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	if has > 0 {
		redisPool.DelKeyByPrefix("article_category_num:")
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpSuccess,
			"msg":    "新增文章成功",
			"url":    "/admin/article/list",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpError,
		"msg":    "新增文章失败",
	})
	return
}

//endregion

//region Remark: 编辑 Author:Qing
func GetArticleEdit(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	item, _ := logic.DefaultArticle.FindOne(c, "id", id)
	if item == nil {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"msg":    "没有该篇文章",
		})
		return
	}
	c.HTML(http.StatusOK, "/article/edit", gin.H{
		"Data":  item,
		"Level": logic.DefaultCategory.LevelList(c),
	})
}

func PostArticleEdit(c *gin.Context) {
	id, _ := strconv.ParseInt(c.PostForm("id"), 10, 64)
	title := c.PostForm("title")
	abstract := c.PostForm("abstract")
	//author := c.PostForm("author")
	cover := c.PostForm("cover")
	tags := c.PostForm("tags")
	sources := c.PostForm("sources")
	content := c.PostForm("content")
	sort, _ := strconv.ParseInt(c.PostForm("sort"), 10, 64)
	category_id, _ := strconv.ParseInt(c.PostForm("category_id"), 10, 64)
	is_show := true
	if c.PostForm("is_show") != "1" {
		is_show = false
	}
	allow_comments := true
	if c.PostForm("allow_comments") != "1" {
		allow_comments = false
	}
	if title == "" {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"msg":    "标题不能为空",
		})
		return
	}
	if cover == "" {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"msg":    "封面图不能为空",
		})
		return
	}
	if content == "" {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"msg":    "内容不能为空",
		})
		return
	}
	item := &model.Article{}
	item.Title = title
	item.Abstract = abstract
	//item.Author = author
	item.Cover = cover
	item.Tags = tags
	item.Sources = sources
	item.IsShow = is_show
	item.Sort = sort
	item.CategoryId = category_id
	item.AllowComments = allow_comments
	item.Content = template.HTML(content)
	//has, err := databases.Orm.Id(id).Cols("title", "abstract", "author", "cover", "tags", "sources", "is_show", "allow_comments", "sort", "category_id", "content").Update(item)
	has, err := databases.Orm.Id(id).Cols("title", "abstract", "cover", "tags", "sources", "is_show", "allow_comments", "sort", "category_id", "content").Update(item)
	if err != nil {
		logger.Errorln("Article controller find errof:", err)
		fmt.Println(err.Error())
	}
	if has > 0 {
		redisPool.DelKeyByPrefix("article_category_num:")
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpSuccess,
			"msg":    "更新成功",
			"url":    "/admin/article/list",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpError,
		"msg":    "更新失败",
	})
	return
}

//endregion

//region Remark: 删除 Author:Qing
func PostArticleDel(c *gin.Context) {
	id, _ := strconv.ParseInt(c.PostForm("id"), 10, 64)
	ids := c.PostFormArray("id[]")

	var has int64
	var err error
	if id == 0 {
		has, err = databases.Orm.In("id", ids).Delete(&model.Article{})
	} else {
		has, err = databases.Orm.Id(id).Delete(&model.Article{})
	}

	if err != nil {
		logger.Errorln("Article controller del a article error")
	}

	if has < 1 {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpSuccess,
			"msg":    "删除失败",
		})
		return
	}
	redisPool.DelKeyByPrefix("article_category_num:")
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpSuccess,
		"msg":    "删除成功",
		"url":    "/admin/article/list",
	})
	return
}

//endregion

type ArticleController struct{}

func (ArticleController) table() string {
	return "article"
}

func (ArticleController) shili() *model.Article {
	return &model.Article{}
}

//region Remark: 删除 Author:Qing
func (self ArticleController) Del(c *gin.Context) {
	id, _ := strconv.ParseInt(c.PostForm("id"), 10, 64)
	ids := c.PostFormArray("id[]")

	var has int64
	var err error
	if id == 0 {
		has, err = databases.Orm.In("id", ids).Delete(self.shili())
	} else {
		has, err = databases.Orm.Id(id).Delete(self.shili())
	}

	if err != nil {
		logger.Errorln("Article controller del a article error")
	}

	if has < 1 {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpSuccess,
			"msg":    "删除失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpSuccess,
		"msg":    "删除成功",
		"url":    "/admin/article/list",
	})
	return
}

//endregion

//region Remark: 排序 Author:Qing
func PostArticleSort(c *gin.Context) {
	sort, _ := strconv.ParseInt(c.PostForm("sort"), 10, 64)
	id, _ := strconv.ParseInt(c.PostForm("id"), 10, 64)
	item := &model.Article{}
	item.Sort = sort
	has, err := databases.Orm.ID(id).Cols("sort").Update(item)
	if has < 1 || err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"msg":    "操作失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpSuccess,
		"msg":    "操作成功",
		"url":    "/admin/article/list",
	})
	return
}

//endregion
