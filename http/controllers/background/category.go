package background

import (
	"github.com/qingcc/go_learn/config"
	"github.com/qingcc/go_learn/databases"
	"fmt"
	"github.com/gin-gonic/gin"
	//"github.com/polaris1119/logger"
	"github.com/polaris1119/logger"
	"github.com/qingcc/go_learn/logic"
	"github.com/qingcc/go_learn/model"
	"net/http"
	"strconv"
)

//region Remark: 列表 Author:Qing
func GetCateList(c *gin.Context) {
	c.HTML(http.StatusOK, "/category/list", gin.H{
		"Title": "BackGround Center",
		"Data":  logic.DefaultCategory.LevelList(c),
	})
}

//endregion

//region Remark: 添加 Author:Qing
func GetCateAdd(c *gin.Context) {
	c.HTML(http.StatusOK, "/category/add", gin.H{
		"Title": "BackGround Center",
		"Level": logic.DefaultCategory.LevelList(c),
	})
}

func PostCateAdd(c *gin.Context) {
	title := c.PostForm("title")
	describe := c.PostForm("describe")
	pid, _ := strconv.ParseInt(c.PostForm("pid"), 10, 64)
	sort, _ := strconv.ParseInt(c.PostForm("sort"), 10, 64)
	is_show := true
	if c.PostForm("is_show") != "true" {
		is_show = false
	}
	//if res, _ := logic.DefaultCategory.FindOne(c, "title", title); res != nil {
	//	c.JSON(http.StatusOK, gin.H{
	//		"status": config.HttpSuccess,
	//		"msg":    "该分类已存在",
	//	})
	//	return
	//}
	has, err := databases.Orm.Insert(&model.Category{Title: title, Describe: describe, Pid: pid, Sort: sort, IsShow: is_show})
	if err != nil {
		fmt.Println(err.Error())
	}
	if has > 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpSuccess,
			"msg":    "新增分类成功",
			"url":    "/admin/category/list",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpError,
		"msg":    "新增分类失败",
	})
	return
}

//endregion

//region Remark: 编辑 Author:Qing
func GetCateEdit(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	item, _ := logic.DefaultCategory.FindOne(c, "id", id)

	c.HTML(http.StatusOK, "/category/edit", gin.H{
		"Data":  item,
		"Level": logic.DefaultCategory.LevelList(c),
	})
}

func PostCateEdit(c *gin.Context) {
	id, _ := strconv.ParseInt(c.PostForm("id"), 10, 64)
	title := c.PostForm("title")
	describe := c.PostForm("describe")
	pid, _ := strconv.ParseInt(c.PostForm("pid"), 10, 64)
	sort, _ := strconv.ParseInt(c.PostForm("sort"), 10, 64)
	is_show := true
	if c.PostForm("is_show") != "true" {
		is_show = false
	}
	if pid == id {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"msg":    "自己不能是自己的父级",
		})
		return
	}

	item := &model.Category{}
	item.Title = title
	item.Describe = describe
	item.Pid = pid
	item.Sort = sort
	item.IsShow = is_show
	has, err := databases.Orm.Id(id).Cols("title", "describe", "pid", "sort", "is_show").Update(item)
	if err != nil {
		logger.Errorln("Category controller find errof:", err)
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

//region Remark: 删除 Author:Qing
func PostCateDel(c *gin.Context) {
	id, _ := strconv.ParseInt(c.PostForm("id"), 10, 64)
	ids := c.PostFormArray("id[]")
	//if logic.DefaultRole.RoleHasAdmin(c, id, ids) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"status": config.HttpError,
	//		"msg":    "分类下有文章,不能删除",
	//	})
	//	return
	//}
	var has int64
	var err error
	if id == 0 {
		has, err = databases.Orm.In("id", ids).Delete(&model.Category{})
	} else {
		has, err = databases.Orm.Id(id).Delete(&model.Category{})
	}

	if err != nil {
		logger.Errorln("Category controller del a role error")
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
		"url":    "/admin/category/list",
	})
	return
}

//endregion

//region Remark: 排序 Author:Qing
func PostCateSort(c *gin.Context) {
	sort, _ := strconv.ParseInt(c.PostForm("sort"), 10, 64)
	id, _ := strconv.ParseInt(c.PostForm("id"), 10, 64)
	banners := new(model.Category)
	banners.Sort = sort
	has, err := databases.Orm.ID(id).Cols("sort").Update(banners)
	if has < 1 || err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"msg":    "操作失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpSuccess,
		"msg":    "操作成功",
		"url":    "/admin/category/list",
	})
	return
}

//endregion
