package model

import (
	"time"
)

type Admin struct {
	Id            int64  `xorm:"pk autoincr int(8)"`
	Email         string `xorm:"varchar(255)"`
	Name          string `xorm:"unique index varchar(255)"`
	Password      string `xorm:"varchar(255)"`
	RoleId        int64  `xorm:"index"`
	LastLoginIp   string `xorm:"varchar(32)"`
	LastLoginTime string `xorm:"varchar(255)"`
	LoginCount    int64  `xorm:"int(8)"`
	IsLock        bool   `xorm:"default false"`
	Describe      string
	CreatedAt     time.Time `xorm:"created"`
	UpdatedAt     time.Time `xorm:"updated"`

	//非表字段
	RoleNodes       []*RoleNode            `xorm:"-"`
	RoleNodesRoutes []*RoleNodeRoutes      `xorm:"-"`
	Role            *Role                  `xorm:"-"`
	NavigationNodes []*AdminNavigationNode `xorm:"-"`
}

type AdminList struct {
	Admin    `xorm:"extends"`
	RoleName string
	Ctime    string
}
