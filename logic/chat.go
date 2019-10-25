package logic

import (
	"github.com/qingcc/goblog/databases"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qingcc/goblog/model"
	"github.com/qingcc/goblog/util"
)

type ChatLogic struct{}

var DefaultChat = ChatLogic{}

//region Remark: 评论 Author:Qing
func (self ChatLogic) GetData(limit int) []*model.Chat {
	data := make([]*model.Chat, 0)
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
func (self ChatLogic) FindOne(c *gin.Context, field string, val interface{}) *model.Chat {
	objLog := GetLogger(c)

	item := &model.Chat{}
	_, err := databases.Orm.Where(field+"= ?", val).Get(item)
	if err != nil {
		objLog.Errorf("RoleLogic find errof:", err)
	}

	return item
}

func (self ChatLogic) FindOneChat(field string, val interface{}) *model.Chat {
	item := &model.Chat{}
	_, err := databases.Orm.Where(field+"= ?", val).Get(item)
	if err != nil {
		fmt.Println("RoleLogic find errof:", err)
	}

	return item
}

//endregion
