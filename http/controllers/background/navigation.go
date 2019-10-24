package background

import (
	"config"
	"databases"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/polaris1119/logger"
	"logic"
	"model"
	"net/http"
	"strconv"
	"util"
)

type SendData struct {
	Name  string
	Other string
}

//region Remark: 导航节点列表 Author:Qing
func GetNavigationList(c *gin.Context) {
	//var data []model.AdminNavigation
	//if res,_ := Exists
	//redis是否有缓存,是,取缓存, 否,查询,加入缓存
	c.HTML(http.StatusOK, "navigation/list", gin.H{
		"title":    "BackGround Center",
		"datalist": logic.DefaultAdminNavigation.DataList(c),
	})
}

//endregion

//region Remark: 添加导航节点 Author:Qing
func GetNavigationAdd(c *gin.Context) {
	msg := &SendData{
		Name:  "测试",
		Other: "23333",
	}
	broadcast <- msg
	c.HTML(http.StatusOK, "navigation/add", gin.H{
		"title":    "BackGround Center",
		"datalist": logic.DefaultAdminNavigation.DataList(c),
	})
}

func PostNavigationAdd(c *gin.Context) {
	title := c.PostForm("title")
	url := c.PostForm("url")
	parent_id, _ := strconv.ParseInt(c.PostForm("parent_id"), 10, 64)
	sort, _ := strconv.ParseInt(c.PostForm("sort"), 10, 64)

	index_max, _ := strconv.Atoi(c.PostForm("nodeids"))
	var is_show = true
	var is_sys = false

	if c.PostForm("is_show") != "true" {
		is_show = false
	}

	if c.PostForm("is_sys") != "false" {
		is_sys = false
	}

	//开启事务
	db := databases.Orm.NewSession()
	defer db.Close()
	err1 := db.Begin()
	if err1 != nil {
		fmt.Println(err1.Error())
	}

	navigation := &model.AdminNavigation{
		Title:    title,
		Url:      url,
		ParentId: parent_id,
		Sort:     sort,
		IsShow:   is_show,
		IsSys:    is_sys,
	}
	_, err := db.Insert(navigation)

	if err != nil {
		db.Rollback()
		fmt.Println(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"msg":    "添加失败",
		})
		return
	}

	if index_max == 0 {
		db.Commit()
		fmt.Println(111)
		//更新 redis 缓存
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpSuccess,
			"msg":    "添加成功",
		})
		return
	}

	//插入节点数据--AdminNovigationNode

	nodes := make([]model.AdminNavigationNode, index_max)
	for i := 0; i < index_max; i++ {
		title := c.PostForm("node[" + strconv.Itoa(i) + "][title]")
		route_action := c.PostForm("node[" + strconv.Itoa(i) + "][route_action]")
		node_sort, _ := strconv.ParseInt(c.PostForm("node["+strconv.Itoa(i)+"][sort]"), 10, 64)
		nodes[i].Title = title
		nodes[i].RouteAction = route_action
		nodes[i].Sort = node_sort
		nodes[i].AdminNavigationId = navigation.Id
	}

	_, err = db.Insert(&nodes)

	if err != nil {
		db.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"msg":    "添加失败",
		})
		return
	} else {
		//db.Rollback()
		db.Commit()
		//更新 redis 缓存
		util.DelKey("admin:navigation:list")
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpSuccess,
			"msg":    "添加成功",
		})
		return
	}
}

//endregion

//region Remark: 编辑导航节点 Author:Qing
func GetNavigationEdit(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	data := logic.DefaultAdminNavigation.Modify(id)

	c.HTML(http.StatusOK, "navigation/edit", gin.H{
		"Title":             "BackGround Center",
		"AllNavigationList": logic.DefaultAdminNavigation.DataList(c),
		"Data":              data,
		"Len":               len(data.Node),
	})
}

func PostNavigationEdit(c *gin.Context) {
	id, _ := strconv.ParseInt(c.PostForm("id"), 10, 64)
	title := c.PostForm("title")
	url := c.PostForm("url")
	parent_id, _ := strconv.ParseInt(c.PostForm("parent_id"), 10, 64)
	sort, _ := strconv.ParseInt(c.PostForm("sort"), 10, 64)

	index_max, _ := strconv.Atoi(c.PostForm("nodeids"))
	index_max = index_max - 1

	var is_show = true
	var is_sys = true

	if c.PostForm("is_show") != "1" {
		is_show = false
	}
	if c.PostForm("is_sys") != "1" {
		is_sys = false
	}

	var redis_key string = "admin:navigation:list"
	//开启事务
	db := databases.Orm.NewSession()
	defer db.Close()
	err1 := db.Begin()
	if err1 != nil {
		fmt.Println(err1.Error())
	}
	data := &model.AdminNavigation{}
	res, err := db.Id(id).Get(data)
	if err != nil {
		db.Rollback()
		fmt.Println(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"msg":    "修改失败",
		})
		return
	}
	if res != true {
		db.Rollback()
		fmt.Println(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"msg":    "修改失败",
		})
		return
	}
	data.Title = title
	data.Url = url
	data.ParentId = parent_id
	data.Sort = sort
	data.IsShow = is_show
	data.IsSys = is_sys
	has, err := db.Id(id).Cols("title", "url", "parent_id", "sort", "is_show", "is_sys").Update(data)
	if err != nil {
		db.Rollback()
		fmt.Println(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"msg":    "修改失败",
		})
		return
	}
	if has < 1 {
		db.Rollback()
		fmt.Println(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"msg":    "修改失败",
		})
		return
	}
	//删除节点数据
	del := &model.AdminNavigationNode{}

	_, err = db.Where("admin_navigation_id = ?", id).Delete(del)
	if err != nil {
		db.Rollback()
		fmt.Println(err.Error())
	}

	if index_max == 0 {
		db.Commit()
		//更新 redis 缓存
		util.DelKey(redis_key)
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpSuccess,
			"msg":    "修改成功",
		})
		return
	}

	//插入节点数据--AdminNovigationNode
	nodes := make([]model.AdminNavigationNode, index_max)
	for i := 0; i < index_max; i++ {
		title := c.PostForm("node[" + strconv.Itoa(i) + "][title]")
		route_action := c.PostForm("node[" + strconv.Itoa(i) + "][route_action]")
		node_sort, _ := strconv.ParseInt(c.PostForm("node["+strconv.Itoa(i)+"][sort]"), 10, 64)
		nodes[i].Title = title
		nodes[i].RouteAction = route_action
		nodes[i].Sort = node_sort
		nodes[i].AdminNavigationId = data.Id
	}

	_, err = db.Insert(&nodes)

	if err != nil {
		db.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"msg":    "修改失败",
		})
		return
	} else {
		//db.Rollback()
		db.Commit()
		//更新 redis 缓存
		util.DelKey(redis_key)
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpSuccess,
			"msg":    "修改成功",
		})
		return
	}
}

//endregion

//region Remark: 导航节点排序 Author; chijian
func PostNavigationSort(c *gin.Context) {
	sort, _ := strconv.ParseInt(c.PostForm("sort"), 10, 64)
	id, _ := strconv.ParseInt(c.PostForm("id"), 10, 64)
	data := logic.DefaultAdminNavigation.Modify(id)
	data.Sort = sort

	has, err := databases.Orm.ID(id).Cols("sort").Update(data)
	if has < 1 || err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "操作失败",
		})
		return
	}
	util.DelKey("admin:navigation:list")
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpSuccess,
		"msg":    "操作成功",
		"url":    "/admin/navigation/list",
	})
	return
}

//endregion

//region Remark: 导航节点删除 Author; chijian
func PostNavigationDel(c *gin.Context) {
	id, _ := strconv.ParseInt(c.PostForm("id"), 10, 64)
	ids := c.PostFormArray("id[]")

	has, err := databases.Orm.In("id", ids).Where("is_sys = true").Count(&model.AdminNavigation{})
	if err != nil {
		logger.Errorln("admin controller del a role error")
	}
	if has > 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"msg":    "系统默认的节点不能删除",
		})
		return
	}

	if id == 0 {
		has, err = databases.Orm.In("id", ids).Where("is_sys = false").Delete(&model.AdminNavigationNode{})
	} else {
		has, err = databases.Orm.Id(id).Where("is_sys = false").Delete(&model.AdminNavigationNode{})
	}

	if err != nil {
		logger.Errorln("admin controller del a role error")
	}

	if has < 1 {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"msg":    "删除失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpSuccess,
		"msg":    "删除成功",
		"url":    "/admin/navigation/list",
	})
	return
}

//endregion
