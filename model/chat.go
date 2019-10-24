package model

import (
	"time"
)

type Chat struct {
	Cid     int64     `json:"cid" xorm:"pk autoincr"`
	Objid   int64     `json:"objid"`
	Objtype int64     `json:"objtype"`
	Content string    `json:"content"`
	Uid     int64     `json:"uid"`
	Floor   int64     `json:"floor"`
	Flag    int64     `json:"flag"`
	Time    time.Time `json:"ctime" xorm:"created"`

	Objinfo    map[string]interface{} `json:"objinfo" xorm:"-"`
	ReplyFloor int64                  `json:"reply_floor" xorm:"-"` // 回复某一楼层
}

func (*Chat) TableName() string {
	return "chat"
}
