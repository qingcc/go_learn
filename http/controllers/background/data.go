package background

import (
	"github.com/qingcc/go_learn/config"
	"github.com/qingcc/go_learn/databases"
	"fmt"
	"github.com/gin-gonic/gin"
	uuid2 "github.com/satori/go.uuid"
	"io/ioutil"
	"github.com/qingcc/go_learn/model"
	"net/http"
	"os"
	"time"
	"github.com/qingcc/go_learn/util"
)

//region Remark: 初始化数据 Author:Qing
func InitData() bool {
	has, err := databases.Orm.Insert(&model.Role{
		RoleName: "超级管理员",
		IsSuper:  true,
		IsSys:    true,
	}, &model.Role{
		RoleName: "普通管理员",
		IsSuper:  false,
		IsSys:    true,
	})
	if err != nil {
		fmt.Println("管理员角色初始化失败" + err.Error())
		return false
	}
	if has < 1 {
		fmt.Println("管理员角色初始化失败")
		return false
	}

	has, err = databases.Orm.Insert(&model.Admin{
		Name:          "admin",
		Email:         "admin@admin.com",
		Password:      "21232f297a57a5a743894a0e4a801fc3", //admin
		RoleId:        1,
		LastLoginIp:   "127.0.0.1",
		LastLoginTime: time.Now().Format("2006-01-02 15:04:05"),
	}, &model.Admin{
		Name:          "qing",
		Email:         "qing@admin.com",
		Password:      "21232f297a57a5a743894a0e4a801fc3", //admin
		RoleId:        2,
		LastLoginIp:   "127.0.0.1",
		LastLoginTime: time.Now().Format("2006-01-02 15:04:05"),
	})
	if err != nil {
		fmt.Println("管理员初始化失败" + err.Error())
		return false
	}
	if has < 1 {
		fmt.Println("管理员初始化失败")
		return false
	}

	return true
}

//endregion

//region Remark: 数据备份列表 Author; chijian
func BackUpList(c *gin.Context) {
	files, _ := ioutil.ReadDir("./uploads/backup/")
	c.HTML(http.StatusOK, "/backup/list", gin.H{
		"title": "BackGround Login",
		"data":  files,
	})
}

//endregion

//region Remark: 备份数据 Author; chijian
func BackUp(c *gin.Context) {
	uuid := uuid2.NewV4()
	file := "./uploads/backup/" + uuid.String() + "_" + time.Now().Format("2006_01_02_15_04_05") + ".sql"
	databases.Orm.DumpAllToFile(file)
	f1, err := os.Open(file)
	fmt.Println(err)
	var files = []*os.File{f1}
	fileZip := "./uploads/backup/" + uuid.String() + "_" + time.Now().Format("2006_01_02_15_04_05") + ".zip"
	util.Zip(files, fileZip)
	os.Remove(file)
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpSuccess,
		"msg":    "备份成功",
		"url":    "/admin/backup/list",
	})
	return

}

//endregion

//region Remark: 下载备份 Author:Qing
func DownloadBackup(c *gin.Context) {
	fileName := "./uploads/backup/" + c.Param("name")
	c.File(fileName)
}

//endregion

//region Remark: 删除备份 Author; chijian
func DelBackup(c *gin.Context) {
	fileName := "./uploads/backup/" + c.Param("name")
	os.Remove(fileName)
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpSuccess,
		"msg":    "删除备份成功",
		"url":    "/admin/backup/list",
	})
	return

}

//endregion
