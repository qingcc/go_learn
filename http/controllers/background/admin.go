package background

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qingcc/goblog/config"
	"github.com/qingcc/goblog/databases"
	//"github.com/polaris1119/logger"
	"github.com/qingcc/goblog/logic"
	"github.com/qingcc/goblog/model"
	"github.com/qingcc/goblog/util"
	"net/http"
	"strconv"
)

//region Remark: 管理员列表 Author:Qing
func GetAdminUserList(c *gin.Context) {

	admin, limit, page, num := logic.DefaultAdmin.GetAdminList(8, 0, "")

	c.HTML(http.StatusOK, "/admin/list", gin.H{
		"title":    "BackGround Login",
		"info":     "",
		"datalist": admin,
		"limit":    limit,
		"page":     page,
		"num":      num,
	})
}

//endregion

//region Remark: 添加管理员 Author:Qing
func GetAdminUserAdd(c *gin.Context) {
	role := logic.DefaultRole.RoleList(c)

	c.HTML(http.StatusOK, "/admin/add", gin.H{
		"title": "BackGround Login",
		"role":  role,
	})
}

func PostAdminUserAdd(c *gin.Context) {
	name := c.PostForm("name")
	password := c.PostForm("password")
	email := c.PostForm("email")
	role_id, _ := strconv.ParseInt(c.PostForm("role_id"), 10, 64)
	describe := c.PostForm("describe")
	if admin := logic.DefaultAdmin.GetAdminByField(c, "email", email); admin.Id > 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpSuccess,
			"msg":    "该邮箱已存在",
		})
		return
	}
	has, err := databases.Orm.Insert(&model.Admin{Name: name, Describe: describe, Email: email, Password: util.Strmd5(password), RoleId: role_id})
	if err != nil {
		fmt.Println(err.Error())
	}
	if has > 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpSuccess,
			"msg":    "新增管理员成功",
			"url":    "/admin/admin_user/list",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpError,
		"msg":    "新增管理员失败",
	})
	return
}

//endregion

//region Remark: 编辑角色 Author:Qing
func GetAdminUserEdit(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	admin := logic.DefaultAdmin.GetAdminByField(c, "id", id)
	role := logic.DefaultRole.RoleList(c)

	c.HTML(http.StatusOK, "/admin/edit", gin.H{
		"data": admin,
		"role": role,
	})
}

func PostAdminUserEdit(c *gin.Context) {
	id, _ := strconv.ParseInt(c.PostForm("id"), 10, 64)
	name := c.PostForm("name")
	email := c.PostForm("email")
	role_id, _ := strconv.ParseInt(c.PostForm("role_id"), 10, 64)
	describe := c.PostForm("describe")
	admin_sour := logic.DefaultAdmin.GetAdminByField(c, "id", id)
	admin_new := logic.DefaultAdmin.GetAdminByField(c, "email", email)

	if admin_sour.Email != email && admin_new.Id > 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpSuccess,
			"msg":    "该邮箱已存在",
		})
		return
	}
	admin := &model.Admin{Name: name, Email: email, RoleId: role_id, Describe: describe}
	has, err := databases.Orm.Id(id).Cols("name", "email", "role_id", "describe").Update(admin)
	if err != nil {
		//logger.Errorln("Admin controller find errof:", err)
		fmt.Println(err.Error())
	}
	if has > 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpSuccess,
			"msg":    "更新成功",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpError,
		"msg":    "更新失败",
	})
	return
}

//endregion

//region Remark: 删除 角色 Author:Qing
func PostAdminUserDel(c *gin.Context) {
	id, _ := strconv.ParseInt(c.PostForm("id"), 10, 64)
	ids := c.PostFormArray("id[]")
	var has int64
	var err error
	if id == 0 {
		has, err = databases.Orm.In("id", ids).Delete(&model.Admin{})
	} else {
		has, err = databases.Orm.Id(id).Delete(&model.Admin{})
	}

	if err != nil {
		//logger.Errorln("admin controller del a role error")
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
		"url":    "/admin/admin_user/list",
	})
	return
}

//endregion

//region Remark: 启用 或 停用 某管理员 Author:Qing
func PostAdminUserEnable(c *gin.Context) {
	id, _ := strconv.ParseInt(c.PostForm("id"), 10, 64)

	admin := &model.Admin{}
	_, err := databases.Orm.Id(id).Get(admin)
	if err != nil {
		//logger.Errorln("admin controller del a role error")
	}
	var isOn = true
	var msg string
	if admin.IsLock {
		admin.IsLock = false
		isOn = false
	} else {
		admin.IsLock = true
	}
	has, err := databases.Orm.Id(id).Cols("is_lock").Update(admin)
	if has < 1 {
		if isOn {
			msg = "停用失败"
		} else {
			msg = "启用失败"
		}
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"msg":    msg,
			"isOn":   isOn,
		})
		return
	} else {
		if isOn {
			msg = "停用成功"
		} else {
			msg = "启用成功"
		}
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpSuccess,
			"msg":    msg,
			"isOn":   isOn,
			"id":     id,
			"url":    "/admin/admin_user/enable",
		})
		return
	}
}

//endregion
