package background

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qingcc/goblog/config"
	"github.com/qingcc/goblog/databases"
	//"github.com/polaris1119/logger"
	"github.com/polaris1119/logger"
	"github.com/qingcc/goblog/logic"
	"github.com/qingcc/goblog/model"
	"github.com/qingcc/goblog/util"
	"net/http"
	"strconv"
)

//region Remark: 管理员角色列表 Author:Qing
func GetRoleList(c *gin.Context) {
	c.HTML(http.StatusOK, "/role/list", gin.H{
		"title":    "BackGround Center",
		"datalist": logic.DefaultRole.RoleList(c),
	})
}

//endregion

//region Remark: 添加新角色 Author:Qing
func GetRoleAdd(c *gin.Context) {
	c.HTML(http.StatusOK, "/role/add", gin.H{
		"Title": "BackGround Center",
		"Data":  logic.DefaultAdminNavigation.DataList(c),
	})
}

func PostRoleAdd(c *gin.Context) {
	role_name := c.PostForm("role_name")
	desccribe := c.PostForm("desccribe")
	if logic.DefaultRole.RoleIsExist(c, role_name) {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpSuccess,
			"msg":    "该角色已存在",
		})
		return
	}

	//开启事物
	db := databases.Orm.NewSession()
	db.Begin()

	role := new(model.Role)
	role.RoleName = role_name
	role.Describe = desccribe
	has, err := db.Insert(role)
	if err != nil {
		db.Rollback()
		fmt.Println(err.Error())
	}
	if has < 0 {
		db.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"msg":    "新增角色失败",
		})
		return
	}
	nodes := c.PostFormArray("node[]")
	adminNavigationNode := make([]model.AdminNavigationNode, 0)
	db.In("id", nodes).Cols("admin_navigation_id").GroupBy("admin_navigation_id").Find(&adminNavigationNode)

	//获取所有权限及其父节点
	var allIds []int64
	for _, value := range adminNavigationNode {
		var ids []int64
		id := logic.DefaultAdminNavigationNode.GetAllNavigationIds(db, value.AdminNavigationId, ids)
		for _, v := range id {
			allIds = append(allIds, v)
		}
	}
	//去重
	allIds = util.RemoveDuplicateAnd0Int64(allIds)
	rolenodes := make([]model.RoleNode, len(allIds))
	//插入节点(是否显示)
	for k, v := range allIds {
		rolenodes[k].RoleId = role.Id
		rolenodes[k].AdminNavigationId = v
	}
	_, err = db.Insert(&rolenodes)
	if err != nil {
		db.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "操作失败" + err.Error(),
		})
		return
	}
	//插入权限路由 (是否有权限)
	nodeRoutes := make([]model.RoleNodeRoutes, len(nodes))
	for key, value := range nodes {
		nodeRoutes[key].AdminNavigationNodeId, _ = strconv.ParseInt(value, 10, 64)
		nodeRoutes[key].RoleId = role.Id
	}
	_, err = db.Insert(nodeRoutes)
	if err != nil {
		db.Rollback()
		fmt.Println(err.Error())
	}
	db.Commit()
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpSuccess,
		"msg":    "新增角色成功",
		"url":    "/admin/admin_role/list",
	})
	return
}

//endregion

//region Remark: 编辑角色 Author:Qing
func GetRoleEdit(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	role, _ := logic.DefaultRole.FindOne(c, "id", id)

	checked := make([]*model.RoleNodeRoutes, 0)
	err := databases.Orm.Where("role_id = ?", id).Find(&checked)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("len(checked):", len(checked))
	checked_ids := ""
	for key, value := range checked {
		if key == 0 {
			checked_ids = strconv.FormatInt(value.AdminNavigationNodeId, 10)
			continue
		}
		checked_ids = checked_ids + "," + strconv.FormatInt(value.AdminNavigationNodeId, 10)
	}
	fmt.Println("ids:", checked_ids)
	c.HTML(http.StatusOK, "/role/edit", gin.H{
		"Title":      "BackGround Center",
		"Data":       logic.DefaultAdminNavigation.DataList(c),
		"Role":       role,
		"CheckedIds": checked_ids,
	})
}

func PostRoleEdit(c *gin.Context) {
	id, _ := strconv.ParseInt(c.PostForm("id"), 10, 64)
	role_name := c.PostForm("role_name")
	desccribe := c.PostForm("describe")
	fmt.Println("has:", c.PostFormArray("node[]"))
	//开启事物
	db := databases.Orm.NewSession()
	db.Begin()

	role, err := logic.DefaultRole.FindOne(c, "id", id)
	if err != nil {
		db.Rollback()
		fmt.Println(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"msg":    "修改角色失败",
		})
		return
	}
	role.RoleName = role_name
	role.Describe = desccribe
	has, err := db.Id(id).Cols("role_name", "describe").Update(role)
	if err != nil {
		db.Rollback()
		fmt.Println(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"msg":    "修改角色失败",
		})
		return
	}
	if has < 0 {
		db.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"msg":    "修改角色失败",
		})
		return
	}

	//删除之后再插入
	//region 删除旧的AdminRoleNode数据
	role_node := new(model.RoleNode)
	_, err = databases.Orm.Where("role_id=?", id).Delete(role_node)
	if err != nil {
		db.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "操作失败" + err.Error(),
		})
		return
	}

	//endregion
	node := c.PostFormArray("node[]")
	admin_navigation_nodes := make([]*model.AdminNavigationNode, 0)
	db.In("id", node).Cols("admin_navigation_id").GroupBy("admin_navigation_id").Find(&admin_navigation_nodes)

	//----获取所有的AdminNavigationId，包括父级的
	var allids []int64
	for _, v := range admin_navigation_nodes {
		var ids []int64
		id := logic.DefaultAdminNavigationNode.GetAllNavigationIds(db, v.AdminNavigationId, ids)
		for _, v1 := range id {
			allids = append(allids, v1)
		}
	}
	allids = util.RemoveDuplicateAnd0Int64(allids)
	rolenodes := make([]*model.RoleNode, 0)
	for _, value := range node {
		navigation_id, _ := strconv.ParseInt(value, 10, 64)
		item := &model.RoleNode{RoleId: role.Id, AdminNavigationId: navigation_id}
		rolenodes = append(rolenodes, item)
	}
	_, err = db.Insert(rolenodes)
	if err != nil {
		db.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"msg":    "修改角色失败",
		})
		return
	}
	//删除原有节点
	_, err = db.Where("role_id = ?", id).Delete(&model.RoleNodeRoutes{})
	if err != nil {
		db.Rollback()
		fmt.Println(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"msg":    "修改角色失败",
		})
		return
	}
	routes := make([]model.RoleNodeRoutes, len(node))
	for key, value := range node {
		routes[key].RoleId = id
		route_id, _ := strconv.ParseInt(value, 10, 64)
		routes[key].AdminNavigationNodeId = route_id
	}
	_, err = db.Insert(routes)
	if err != nil {
		db.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"msg":    "修改角色失败",
		})
		return
	}

	db.Commit()

	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpSuccess,
		"msg":    "修改成功",
	})
	return

}

//endregion

//region Remark: 删除 角色 Author:Qing
func PostRoleDel(c *gin.Context) {
	id, _ := strconv.ParseInt(c.PostForm("id"), 10, 64)
	ids := c.PostFormArray("id[]")
	if logic.DefaultRole.RoleHasAdmin(c, id, ids) {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"msg":    "角色下有管理员,不能删除",
		})
		return
	}
	var has int64
	var err error
	if id == 0 {
		has, err = databases.Orm.In("id", ids).Delete(&model.Role{})
	} else {
		has, err = databases.Orm.Id(id).Delete(&model.Role{})
	}

	if err != nil {
		logger.Errorln("role controller del a role error")
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
		"url":    "/admin/admin_role/list",
	})
	return
}

//endregion
