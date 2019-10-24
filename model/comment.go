package model

import (
	"time"
)

// 评论信息（通用）
type Comment struct {
	Cid     int       `json:"cid" xorm:"pk autoincr"`
	Objid   int       `json:"objid"`
	Objtype int       `json:"objtype"`
	Content string    `json:"content"`
	Uid     int       `json:"uid"`
	Floor   int       `json:"floor"`
	Flag    int       `json:"flag"`
	Time    time.Time `json:"ctime" xorm:"created"`

	Objinfo    map[string]interface{} `json:"objinfo" xorm:"-"`
	ReplyFloor int                    `json:"reply_floor" xorm:"-"` // 回复某一楼层
}

func (*Comment) TableName() string {
	return "comments"
}
