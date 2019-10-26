package logic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qingcc/goblog/databases"
	"github.com/qingcc/goblog/model"
)

type RoleNodeRoutesLogic struct{}

var DefaultRoleNodeRoutes = RoleNodeRoutesLogic{}

//region Remark: 获取角色的权限 Author:Qing
func (self RoleNodeRoutesLogic) GetRoleNodeRoutes(c *gin.Context, role_id int64) []*model.RoleNodeRoutes {
	objLog := GetLogger(c)

	nodes := make([]*model.RoleNodeRoutesAndRouteAction, 0)
	err := databases.Orm.Table("role_node_routes").Where("role_id = ?", role_id).
		Join("LEFT", "admin_navigation_node", "admin_navigation_node.id = role_node_routes.admin_navigation_node_id").Find(&nodes)
	if err != nil {
		objLog.Errorf("RoleNodeRoutesLogic find errof:", err)
		return nil
	}
	data := make([]*model.RoleNodeRoutes, len(nodes))
	for key, value := range nodes {
		value.RoleNodeRoutes.RouteAction = nodes[key].AdminNavigationNode.RouteAction
		data[key] = &value.RoleNodeRoutes
	}

	return data
}

//endregion

//region Remark: 判断是否具有该权限 Author:Qing
func (self RoleNodeRoutesLogic) HasAuth(routes []*model.RoleNodeRoutes, routename string) bool {
	fmt.Println(111)
	for _, value := range routes {
		if value.RouteAction == routename {
			return true
		}
	}
	return false
}

//endregion

//region Remark: 判断是否具有该权限 Author:Qing
func (self RoleNodeRoutesLogic) FindOne(c *gin.Context, field string, val interface{}) *model.RoleNodeRoutes {
	objLog := GetLogger(c)
	route := &model.RoleNodeRoutes{}
	res, err := databases.Orm.Where(field+"= ?", val).Get(route)
	if err != nil {
		objLog.Errorf("RoleNodeRoutesLogic find errof:", err)
		return nil
	}
	if res == false {
		return nil
	}
	return route
}

//endregion
