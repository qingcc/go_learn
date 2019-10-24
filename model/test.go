package model

type Test struct {
	Id     int64   `xorm:"pk autoincr BIGINT"  json:"id"`
	CoinId int64   `xorm:"not null int" json:"coin_id"` //货币id
	High   float64 `xorm:"not null decimal(12, 2)"`     //最高价
	Low    float64 `xorm:"not null Numeric"`            //最低价
}
