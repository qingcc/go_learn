package logic

import (
	"github.com/qingcc/go_learn/databases"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/qingcc/go_learn/model"
	"strings"
	"github.com/qingcc/go_learn/util"
)

type AdminNavigationLogic struct{}

var DefaultAdminNavigation = AdminNavigationLogic{}

//region Remark: 权限节点列表 Author:Qing
func (self AdminNavigationLogic) DataList(c *gin.Context) []*model.AdminNavigation {
	objLog := GetLogger(c)
	navigation := make([]*model.AdminNavigation, 0)

	var redis_key string = "admin:navigation:list"
	if res, _ := util.Exists(redis_key); res == true {
		valueBytes, _ := redis.Bytes(util.Get(redis_key))
		json.Unmarshal(valueBytes, &navigation)
		return navigation
	} else {
		fmt.Println("do not have session")
		err := databases.Orm.Asc("parent_id").Asc("sort").Find(&navigation)
		if err != nil {
			fmt.Println(err.Error())
			objLog.Errorf("RoleLogic find errof:", err)
			return nil
		}

		showList := make([]*model.AdminNavigation, 0, len(navigation))
		self.Data2List(&showList, navigation, 0, 1, 0, "|-")

		//缓存到redis
		fmt.Println(showList)
		value, _ := json.Marshal(showList)
		util.Set(redis_key, value, 60*60)
		return showList
	}
}

//endregion

//region Remark: 将数据排序 Author:Qing
func (self AdminNavigationLogic) Data2List(showNodeList *[]*model.AdminNavigation, nodeList []*model.AdminNavigation, parentId, curLevel, pos int64, html string) {
	for num := len(nodeList); pos < int64(num); pos++ {
		node := nodeList[pos]
		if node.ParentId == parentId {
			newhtml := strings.Repeat(html, int(curLevel-1))
			node.TitleHtml = newhtml + node.Title

			//region Remark: 拼接 操作 admin_navigation_node Author:Qing
			action := new([]*model.AdminNavigationNode)
			err := databases.Orm.Where("admin_navigation_id = ?", node.Id).Find(action)
			if err != nil {
				fmt.Println("AdminNavigationLogic Data2List search nodes error")
			}
			if len(*action) > 0 {
				str := ""
				for _, value := range *action {
					str = str + "," + value.Title
				}
				node.Action = node.Action + "[" + beego.Substr(str, 1, len(str)) + "]"
				node.Node = *action
			}
			//endregion
			*showNodeList = append(*showNodeList, node)
			self.Data2List(showNodeList, nodeList, node.Id, curLevel+1, 0, html)
		}
	}
}

//endregion

//region Remark: 获取编辑某一导航节点 Author:Qing
func (self AdminNavigationLogic) Modify(id int64) *model.AdminNavigation {
	navigation := &model.AdminNavigation{}
	ok, _ := databases.Orm.Id(id).Get(navigation)

	if !ok {
		fmt.Println("无该导航节点")
		return &model.AdminNavigation{}
	}
	node := make([]*model.AdminNavigationNode, 0)
	err := databases.Orm.Where("admin_navigation_id = ?", navigation.Id).Find(&node)
	if err != nil {
		fmt.Println("AdminNavigationLogic modify error")
		return &model.AdminNavigation{}
	}

	navigation.Node = node

	return navigation
}

//endregion
