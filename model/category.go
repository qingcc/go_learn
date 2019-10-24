package model

import (
	"time"
)

type Category struct {
	Id        int64     `xorm:"pk autoincr int(8)"`
	Title     string    `xorm:"unique varchar(255)"`
	IsShow    bool      `xorm:"default true"`
	Sort      int64     `xorm:"not null int"`
	Describe  string    `xorm:"varchar(255)"`
	Pid       int64     `xorm:"default 0 int"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`

	//非表字段
	TitleHtml string `xorm:"-"`
	TimeHtml  string `xorm:"-"`
}
