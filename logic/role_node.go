package logic

import (
	"blog_go/databases"
	"github.com/gin-gonic/gin"
	"blog_go/model"
)

type RoleNodeLogic struct{}

var DefaultRoleNode = RoleNodeLogic{}

//region Remark: 获取角色 Author:Qing
func (self RoleNodeLogic) FindRoleNodes(c *gin.Context, role_id int64) []*model.RoleNode {
	objLog := GetLogger(c)
	nodes := make([]*model.RoleNode, 0)
	err := databases.Orm.Where("role_id = ?", role_id).Find(&nodes)
	if err != nil {
		objLog.Errorf("RoleNodeLogic find errof:", err)
		return nil
	}

	return nodes
}

//endregion
