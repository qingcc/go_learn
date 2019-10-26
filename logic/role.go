package logic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qingcc/goblog/databases"
	"github.com/qingcc/goblog/model"
)

type RoleLogic struct{}

var DefaultRole = RoleLogic{}

//region Remark: 列表 Author:Qing
func (self RoleLogic) RoleList(c *gin.Context) []*model.Role {
	objLog := GetLogger(c)
	role := make([]*model.Role, 0)
	err := databases.Orm.Find(&role)
	if err != nil {
		objLog.Errorf("RoleLogic find errof:", err)
		return nil
	}
	return role
}

//endregion

//region Remark: 角色是否存在 Author:Qing
func (self RoleLogic) RoleIsExist(c *gin.Context, name string) bool {
	objLog := GetLogger(c)
	has, err := databases.Orm.Where("role_name = ?", name).Count(&model.Role{})
	if err != nil {
		objLog.Errorf("RoleLogic find errof:", err)
	}
	if has == 0 {
		return false
	}
	return true
}

//endregion

//region Remark: 获取角色 Author:Qing
func (self RoleLogic) FindOne(c *gin.Context, field string, val interface{}) (*model.Role, error) {
	objLog := GetLogger(c)

	role := &model.Role{}
	_, err := databases.Orm.Where(field+" = ?", val).Get(role)
	if err != nil {
		objLog.Errorf("RoleLogic find errof:", err)
	}

	return role, err
}

//endregion

//region Remark: 角色下是否有管理员 Author:Qing
func (self RoleLogic) RoleHasAdmin(c *gin.Context, id int64, ids []string) bool {
	objLog := GetLogger(c)

	has1, err := databases.Orm.Table("admin").In("role_id", ids).Count()
	has2, err := databases.Orm.Table("admin").Where("role_id = ?", id).Count()
	if err != nil {
		objLog.Errorf("AdminLogic find errof:", err)
	}
	fmt.Println(has1, has2)
	if has1 > 0 || has2 > 0 {
		return true
	}
	return false
}

//endregion
