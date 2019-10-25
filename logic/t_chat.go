package logic

import (
	"blog_go/databases"
	"github.com/gin-gonic/gin"
	"blog_go/model"
	"blog_go/util"
)

type TChatLogic struct{}

var DefaultTChat = ChatLogic{}

//region Remark: 评论 Author:Qing
func (self TChatLogic) GetData(limit int) []*model.TChat {
	data := make([]*model.TChat, 0)
	count, _ := databases.Orm.Count(&model.Chat{})
	_count, start := int(count), 0
	if _count > limit {
		start = _count - 10
	}

	err := databases.Orm.Asc("time").Limit(limit, start).Find(&data)
	util.CheckErr(err)
	return data
}

//endregion

// region Remark: 获取1条数据 Author:Qing
func (self TChatLogic) FindOne(c *gin.Context, field string, val interface{}) *model.TChat {
	objLog := GetLogger(c)

	item := &model.TChat{}
	_, err := databases.Orm.Where(field+"= ?", val).Get(item)
	if err != nil {
		objLog.Errorf("RoleLogic find errof:", err)
	}

	return item
}

//endregion
