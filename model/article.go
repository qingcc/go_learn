package model

import (
	"html/template"
	"time"
)

type Article struct {
	Id            int64         `xorm:"pk autoincr int(8)"`
	Title         string        `xorm:"varchar(255)"`      //标题
	IsShow        bool          `xorm:"default true"`      //是否展示
	Sort          int64         `xorm:"not null int"`      //排序
	Abstract      string        `xorm:"varchar(255)"`      //简介
	Author        string        `xorm:"varchar(255)"`      //作者
	Content       template.HTML `xorm:"text"`              //文章内容
	Cover         string        `xorm:"varchar(255)"`      //封面图
	CategoryId    int64         `xorm:"default 0 int"`     //分类id
	Tags          string        `xorm:"varchar(255)"`      //标签
	Sources       string        `xorm:"varchar(255)"`      //来源
	AllowComments bool          `xorm:"bool default true"` //允许评论
	Top           int64         `xorm:"tinyint"`           //是否置顶
	Status        int64         `xorm:"tinyint default 1"` //状态
	ViewNum       int64         `xorm:"default 0 int(8)"`  //浏览量
	CommentNum    int64         `xorm:"default 0 int(8)"`  //评论量
	LikesNum      int64         `xorm:"default 0 int(8)"`
	CreatedAt     time.Time     `xorm:"created"`
	UpdatedAt     time.Time     `xorm:"updated"`

	//非表字段
	CategoryHtml string `xorm:"-"`
	TimeHtml     string `xorm:"-"`
}

func (*Article) table() string {
	return "article"
}
