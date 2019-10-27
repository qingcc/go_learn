package model

import (
	"time"
)

type Category struct {
	Id        int64     `xorm:"pk autoincr int(8)" json:"id"`
	Title     string    `xorm:"unique varchar(255)" json:"title"`
	IsShow    bool      `xorm:"default true" json:"is_show"`
	Sort      int64     `xorm:"not null int" json:"sort"`
	Describe  string    `xorm:"varchar(255)" json:"describe"`
	Pid       int64     `xorm:"default 0 int" json:"pid"`
	CreatedAt time.Time `xorm:"created" json:"created_at"`
	UpdatedAt time.Time `xorm:"updated" json:"updated_at"`

	//非表字段
	TitleHtml string `xorm:"-" json:"title_html"`
	TimeHtml  string `xorm:"-" json:"time_html"`
}
