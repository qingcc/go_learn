package model

import (
	"time"
)

type Role struct {
	Id        int64     `xorm:"pk autoincr int(8)"`
	RoleName  string    `xorm:"unique varchar(255)"`
	IsSuper   bool      `xorm:"default false"`
	IsSys     bool      `xorm:"default true"`
	Describe  string    `xorm:"varchar(255)"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}

type RoleNode struct {
	Id                int64     `xorm:"pk autoincr int(8)"`
	RoleId            int64     `xorm:"int(8) not null"`
	AdminNavigationId int64     `xorm:"int(8) not null"`
	CreatedAt         time.Time `xorm:"created"`
	UpdatedAt         time.Time `xorm:"updated"`
}

type RoleNodeRoutes struct {
	Id                    int64     `xorm:"pk autoincr int(8)"`
	RoleId                int64     `xorm:"int(8) not null"`
	AdminNavigationNodeId int64     `xorm:"int(8) not null"`
	CreatedAt             time.Time `xorm:"created"`
	UpdatedAt             time.Time `xorm:"updated"`

	//非表字段
	RouteAction string `xorm:"-"`
}

type RoleNodeRoutesAndRouteAction struct {
	RoleNodeRoutes      `xorm:"extends"`
	AdminNavigationNode `xorm:"extends"`
}
