package logic

import (
	"blog_go/databases"
	"fmt"
	"github.com/gin-gonic/gin"
	"blog_go/model"
	"strings"
)

type CategoryLogic struct{}

var DefaultCategory = CategoryLogic{}

//region Remark: 列表 Author:Qing
func (self CategoryLogic) List(c *gin.Context) []*model.Category {
	objLog := GetLogger(c)
	data := make([]*model.Category, 0)
	err := databases.Orm.Find(&data)
	if err != nil {
		objLog.Errorf("CategoryLogic find errof:", err)
		return nil
	}
	return data
}

//endregion

//region Remark: 获取分类列表 Author:Qing
func (self CategoryLogic) LevelList(c *gin.Context) []*model.Category {
	objLog := GetLogger(c)
	data := make([]*model.Category, 0)
	err := databases.Orm.Asc("pid").Asc("sort").Find(&data)
	if err != nil {
		objLog.Errorf("CategoryLogic find errof:", err)
	}

	showList := make([]*model.Category, 0, len(data))
	self.Data2List(&showList, data, 0, 1, 0, "|-")
	return showList
}

//endregion

//region Remark: 将数据排序 Author:Qing
func (self CategoryLogic) Data2List(showNodeList *[]*model.Category, dataList []*model.Category, parentId, curLevel, pos int64, html string) {
	for num := len(dataList); pos < int64(num); pos++ {
		item := dataList[pos]
		if item.Pid == parentId {
			newhtml := strings.Repeat(html, int(curLevel-1))
			item.TitleHtml = newhtml + item.Title
			item.TimeHtml = item.CreatedAt.Format("2006-01-02 15:04:05")
			*showNodeList = append(*showNodeList, item)
			self.Data2List(showNodeList, dataList, item.Id, curLevel+1, 0, html)
		}
	}
}

//endregion

//region Remark: 获取角色 Author:Qing
func (self CategoryLogic) FindOne(c *gin.Context, field string, val interface{}) (*model.Category, error) {
	objLog := GetLogger(c)

	item := &model.Category{}
	ok, err := databases.Orm.Where(field+" = ?", val).Get(item)
	if err != nil {
		objLog.Errorf("CategoryLogic find errof:", err)
	}
	if ok != true {
		return nil, err
	}
	return item, err
}

//endregion

//region Remark: 角色下是否有管理员 Author:Qing
func (self CategoryLogic) RoleHasAdmin(c *gin.Context, id int64, ids []string) bool {
	objLog := GetLogger(c)

	has1, err := databases.Orm.Table("admin").In("role_id", ids).Count()
	has2, err := databases.Orm.Table("admin").Where("role_id = ?", id).Count()
	if err != nil {
		objLog.Errorf("CategoryLogic find errof:", err)
	}
	fmt.Println(has1, has2)
	if has1 > 0 || has2 > 0 {
		return true
	}
	return false
}

//endregion
