package model

import (
	"time"
)

type AdminNavigation struct {
	Id        int64     `json:"id" xorm:"pk autoincr int(8)"`
	Url       string    `json:"url" xorm:"'url' varchar(255)"`
	Title     string    `json:"title" xorm:"'title' index varchar(255)"`
	ParentId  int64     `json:"parent_id" xorm:"'parent_id' int(8)"`
	IsShow    bool      `json:"is_show" xorm:"'is_show' bool default(true)"`
	IsSys     bool      `json:"is_sys" xorm:"'is_sys' bool default(true)"`
	Sort      int64     `json:"sort" xorm:"'sort' int(8)"`
	CreatedAt time.Time `json:"created_at" xorm:"'created_at' created"`
	UpdatedAt time.Time `json:"updated_at" xorm:"updated"`

	Node      []*AdminNavigationNode `xorm:"-"`
	Level     int64                  `xorm:"-"`
	TitleHtml string                 `xorm:"-"`
	Action    string                 `xorm:"-"`
	HasAction bool                   `xorm:"-"`
}
