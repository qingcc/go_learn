package logic

import (
	"databases"
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"model"
	"strconv"
	"util"
)

type AdminLogic struct{}

var DefaultAdmin = AdminLogic{}

//region Remark: 通过字段查询 管理员表 获取一个实例 Author:Qing
func (self AdminLogic) GetAdminByField(c *gin.Context, field string, val interface{}) *model.Admin {
	objLog := GetLogger(c)
	admin := &model.Admin{}
	_, err := databases.Orm.Where(field+"=?", val).Get(admin)
	if err != nil {
		objLog.Errorf("AdminLogic find errof:", err)
	}
	return admin
}

//endregion

//region Remark: 管理员列表 Author:Qing
func (self AdminLogic) GetAdminList(limit, page int, keyword string) ([]*model.AdminList, int, int, int) {
	admin := make([]*model.AdminList, 0)
	err := databases.Orm.Table("admin").Join("INNER", "role", "admin.role_id = role.id")
	if keyword != "" {
		err = err.Where("name = ?", keyword)
	}

	err1 := *err
	num, _ := err1.Count()

	err2 := err.Limit(limit, limit*page).Find(&admin)
	if err2 != nil {
		fmt.Println(err2.Error())
	}

	return admin, limit, page + 1, int(num)
}

//endregion

/**
当前登陆管理员的信息
*/
func (self AdminLogic) AdminNowLoginInfo(c *gin.Context) *model.Admin {
	var admin *model.Admin
	var adminid int64 = (util.GetSession(c, "adminid")).(int64)
	//缓存数据到redis
	redis_key := "admin:info:" + strconv.FormatInt(adminid, 10)
	if res, _ := util.Exists(redis_key); res == true {
		valueBytes, _ := redis.Bytes(util.Get(redis_key))
		json.Unmarshal(valueBytes, &admin)
	} else {
		admin = DefaultAdmin.GetAdminByField(c, "id", adminid)
		role, _ := DefaultRole.FindOne(c, "id", admin.RoleId)
		admin.Role = role

		admin.RoleNodes = DefaultRoleNode.FindRoleNodes(c, admin.RoleId)
		admin.RoleNodesRoutes = DefaultRoleNodeRoutes.GetRoleNodeRoutes(c, admin.RoleId)
		//缓存到redis
		value, _ := json.Marshal(admin)
		util.Set(redis_key, value, 60*60)
	}
	return admin
}
