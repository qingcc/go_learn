package util

import (
	"github.com/qingcc/go_learn/config"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"io"
	"github.com/qingcc/go_learn/logic"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
	"github.com/qingcc/go_learn/util"
)

//region Remark: 上传 Author:Qing
func UploadImage(c *gin.Context) {
	c.JSON(http.StatusOK, upload(c, "images", "bmp,gif,jpg,jpeg,jpe,png"))
}

//endregion

//region Remark: 上传 Author:Qing
func UploadFile(c *gin.Context) {
	c.JSON(http.StatusOK, upload(c, "file", "zip,rar,pdf,apk"))
}

//endregion

//region Remark: 上传 Author:Qing
func UploadVideo(c *gin.Context) {
	c.JSON(http.StatusOK, upload(c, "video", "mp4"))
}

//endregion

//region Remark: 保存上传的文件 Author:Qing
func upload(c *gin.Context, fileType string, suffix string) gin.H {
	objLog := logic.GetLogger(c)

	//得到上传的文件
	file, header, err := c.Request.FormFile("FileData") //image这个是uplaodify参数定义中的   'fileObjName':'image'
	if err != nil {
		return gin.H{
			"status": config.HttpError,
			"name":   header.Filename,
			"msg":    "上传出现错误",
			"size":   header.Size,
			"data":   "/",
			"url":    "",
		}
	}

	filename := strings.Split(header.Filename, ".")
	filename_suffix := filename[len(filename)-1]
	uuid := uuid.NewV4()
	new_filename := uuid.String() + "." + filename_suffix

	//判断文件后缀是否允许上传
	if !strings.Contains(suffix, filename_suffix) {
		return gin.H{
			//"status": strconv.Itoa(config.HttpError),
			"status": config.HttpError,
			"name":   header.Filename,
			"msg":    "上传格式不允许，只允许上传上传：" + suffix,
			"size":   header.Size,
			"data":   "/",
			"url":    "",
		}
	}

	//创建文件夹
	path := "uploads/" + fileType + "/" + time.Now().Format("2006/0102/")
	util.DirectoryMkdir(path)

	//创建文件
	out, err := os.Create(path + new_filename)
	if err != nil {
		_, file, line, _ := runtime.Caller(0) //获取错误文件和错误行
		objLog.Errorf(file+":"+strconv.Itoa(line), "上传错误：%s", err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		_, file, line, _ := runtime.Caller(0) //获取错误文件和错误行
		objLog.Errorf(file+":"+strconv.Itoa(line), "上传错误：%s", err)
	}
	imgHost := "http://" + c.Request.Host
	//返回值
	return gin.H{
		"status": config.HttpError,
		"name":   header.Filename,
		"msg":    "上传成功",
		"size":   header.Size,
		"data":   "/" + path + new_filename,
		"url":    imgHost + "/" + path + new_filename,
		"path":   "/" + path + new_filename,
	}
}

//endregion
