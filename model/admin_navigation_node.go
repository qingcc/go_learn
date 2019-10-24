package model

import (
	"time"
)

type AdminNavigationNode struct {
	Id                int64     `xorm:"pk autoincr int(8)"`
	AdminNavigationId int64     `xorm:"int(8)"`
	RouteAction       string    `xorm:"varchar(255)"`
	Title             string    `xorm:"varchar(255)"`
	Sort              int64     `xorm:"int(8)"`
	CreatedAt         time.Time `xorm:"created"`
	UpdatedAt         time.Time `xorm:"updated"`
}
