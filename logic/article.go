package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/qingcc/goblog/databases"
	"github.com/qingcc/goblog/model"
	"strings"
)

type ArticleLogic struct{}

var DefaultArticle = ArticleLogic{}

//region Remark: 列表 Author:Qing
type ArticleList struct {
	model.Article  `xorm:"extends"`
	model.Category `xorm:"extends"`
}

func (self ArticleLogic) List(c *gin.Context) []*model.Article {
	objLog := GetLogger(c)
	data := make([]*ArticleList, 0)
	err := databases.Orm.Table("article").Join("LEFT", "category", "article.category_id = category.id").Find(&data)
	if err != nil {
		objLog.Errorf("CategoryLogic find errof:", err)
		return nil
	}
	list := make([]*model.Article, len(data))
	for key, value := range data {
		item := value.Article
		item.CategoryHtml = value.Category.Title
		item.TimeHtml = value.Article.CreatedAt.Format("2006-01-02 15:04:05")
		list[key] = &item
	}
	return list
}

//endregion

func (self ArticleLogic) All(c *gin.Context) []map[string]interface{} {
	objLog := GetLogger(c)
	data := make([]*model.Article, 0)
	err := databases.Orm.Table("article").Desc("id").Find(&data)
	if err != nil {
		objLog.Errorf("CategoryLogic find errof:", err)
		return nil
	}
	list := make([]map[string]interface{}, len(data))
	for key, value := range data {
		item := make(map[string]interface{})
		item["title"] = value.Title
		item["time"] = value.CreatedAt.Format("2006-01-02 15:04")
		item["img"] = value.Cover
		item["abstract"] = value.Abstract
		item["tags"] = strings.Split(value.Tags, ",")
		list[key] = item
	}
	return list
}

//region Remark: 获取一条数据 Author:Qing
func (self ArticleLogic) FindOne(c *gin.Context, field string, val interface{}) (*model.Article, error) {
	objLog := GetLogger(c)

	item := &model.Article{}
	ok, err := databases.Orm.Where(field+" = ?", val).Get(item)
	if err != nil {
		objLog.Errorf("CategoryLogic find errof:", err)
	}
	if !ok {
		return nil, err
	}
	return item, err
}

//endregion
