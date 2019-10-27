package logic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

type InstallLogic struct{}

var DefaultInstall = InstallLogic{}

func (InstallLogic) Createtable(c *gin.Context) {
	objlog := GetLogger(c)
	file_string := "config/db.sql"
	buf, err := ioutil.ReadFile(file_string)
	if err != nil {
		objlog.Errorf("读取db.sql文件失败")
	}

	fmt.Println(string(buf))
}
