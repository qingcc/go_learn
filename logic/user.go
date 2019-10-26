package logic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qingcc/goblog/databases"
	"github.com/qingcc/goblog/model"
)

type UserLogic struct{}

var DefaultUser = UserLogic{}

//region Remark: 列表 Author:Qing
func (self RoleLogic) UserList(c *gin.Context) []*model.User {
	objLog := GetLogger(c)
	item := make([]*model.User, 0)
	err := databases.Orm.Find(&item)
	if err != nil {
		objLog.Errorf("RoleLogic find errof:", err)
		return nil
	}
	return item
}

//endregion

//region Remark: 获取1条数据 Author:Qing
func (self UserLogic) FindOne(c *gin.Context, field string, val interface{}) *model.User {
	objLog := GetLogger(c)

	item := &model.User{}
	_, err := databases.Orm.Where(field+" = ?", val).Get(item)

	if err != nil {
		objLog.Errorf("RoleLogic find errof:", err)
	}

	return item
}

func (self UserLogic) GetOne(field string, val interface{}) *model.User {
	item := &model.User{}
	_, err := databases.Orm.Where(field+" = ?", val).Get(item)

	if err != nil {
		fmt.Println("RoleLogic find errof:", err)
	}

	return item
}

//endregion

//region Remark: 获取1条数据 Author:Qing
func (self UserLogic) FindOneUser(c *gin.Context, username string) *model.User {
	objLog := GetLogger(c)

	item := &model.User{}
	_, err := databases.Orm.Where("username = ? or email = ?", username, username).Get(item)
	if err != nil {
		objLog.Errorf("RoleLogic find errof:", err)
	}

	return item
}

//endregion

// region Remark: 获取1条数据 Author:Qing
func (self UserLogic) FindIn(c *gin.Context, ids []string) []*model.User {
	objLog := GetLogger(c)

	data := make([]*model.User, 0)
	err := databases.Orm.In("id", ids).Find(&data)
	if err != nil {
		objLog.Errorf("RoleLogic find errof:", err)
	}

	return data
}

//endregion
