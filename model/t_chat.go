package model

import (
	"time"
)

type TChat struct {
	Id      int64     `json:"id" xorm:"pk autoincr"`
	Content string    `json:"content"`
	Uid     int64     `json:"uid"`
	Time    time.Time `json:"time" xorm:"created"`
	ToUid   int64     `json:"to_uid"`
}

func (*TChat) TableName() string {
	return "t_chat"
}
