package logic

import (
	//"github.com/polaris1119/logger"
	"github.com/qingcc/goblog/model"
)

var (
	Roles []*model.Role
)

//region Remark: 将所有 角色 加载到内存中；后台修改角色时，重新加载一次 Author:Qing
//func LoadRoles() error {
//	roles := make([]*model.Role, 0)
//	err := databases.Orm.Find(&roles)
//	if err != nil {
//		logger.Errorln("LoadRolers role read failed:", err)
//	}
//
//	if len(roles) < 1 {
//		logger.Errorln("LoadRolers role read failed:num is 0")
//		return errors.New("no role")
//	}
//
//	//maxRoleId := roles[len(roles) - 1].Id
//	Roles = make([]*model.Role, len(roles))
//
//	// 由于角色不多，而且一般角色id是连续自增的，因此这里以角色id当slice的index
//	for _, value := range roles {
//		Roles[value.Id-1] = value
//	}
//
//	logger.Infoln("LoadRoles successfully!")
//
//	return nil
//}

//endregion

//region Remark: 将所有 角色拥有的权限 加载到内存中；后台修改时，重新加载一次 Author:Qing
func LoadRoleAuthorities() error {
	//roleAuthorities := make([]*)
	return nil
}

//endregion
