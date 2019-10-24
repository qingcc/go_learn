package model

import (
	"html/template"
	"time"
)

type Test1 struct {
	Id      int64   `xorm:"pk autoincr BIGINT"  json:"id"`
	CoinId  int64   `xorm:"not null int" json:"coin_id"` //货币id
	High    float64 //最高价
	Title   string
	Content template.HTML `xorm:"text"`
	IsSys   bool
	Tim     time.Time
}
