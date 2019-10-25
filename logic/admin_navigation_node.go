package logic

import (
	"github.com/qingcc/blog_go/databases"
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	"github.com/qingcc/blog_go/model"
)

type AdminNavigationNodeLogic struct{}

var DefaultAdminNavigationNode = AdminNavigationNodeLogic{}

//region Remark: 列表 Author:Qing
func (self AdminNavigationNodeLogic) GetAllNavigationIds(db *xorm.Session, id int64, ids []int64) []int64 {
	admin_navigation := new(model.AdminNavigation)
	db.Id(id).Get(admin_navigation)
	if admin_navigation.ParentId > 0 {
		ids = append(ids, admin_navigation.Id)
		return self.GetAllNavigationIds(db, admin_navigation.ParentId, ids)
	} else {
		ids = append(ids, admin_navigation.Id)
		return ids
	}
}

//endregion

//region Remark: 通过字段查询  获取一个实例 Author:Qing
func (self AdminNavigationNodeLogic) FindOne(c *gin.Context, field string, val interface{}) *model.AdminNavigationNode {
	objLog := GetLogger(c)
	item := &model.AdminNavigationNode{}
	_, err := databases.Orm.Where(field+" = ?", val).Get(item)
	if err != nil {
		objLog.Errorf("AdminLogic find errof:", err)
		return nil
	}
	return item
}

//endregion

////region Remark: 获取角色的权限 Author:Qing
//func (self AdminNavigationNodeLogic) GetNavigationNodes(c *gin.Context, role_id int64) []*model.AdminNavigationNode {
//	objLog := GetLogger(c)
//
//	nodes := make([]*model.AdminNavigationNode, 0)
//	err := databases.Orm.Where("id in (?)", role_id).Find(&nodes)
//	if err != nil {
//		objLog.Errorf("AdminNavigationNodeLogic find errof:", err)
//		return nil
//	}
//
//	return nodes
//}
//
////endregion
