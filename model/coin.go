package model

import (
	"blog_go/databases"
	"fmt"
	"time"
)

type Coin struct {
	Id        int64     `xorm:"pk autoincr BIGINT"  json:"id"`
	Name      string    `xorm:"not null unique VARCHAR(255)" json:"name"`
	Code      string    `xorm:"not null unique VARCHAR(255)" json:"code"`
	Account   string    `xorm:"not null VARCHAR(255)" json:"account"`
	Pwd       string    `xorm:"not null VARCHAR(255)" json:"pwd"`
	AllowIp   string    `xorm:"not null VARCHAR(255)" json:"allow_ip"`
	Port      string    `xorm:"not null VARCHAR(255)" json:"port"`
	Img       string    `xorm:"not null VARCHAR(255)" json:"img"`
	IsTrade   bool      `xorm:"not null default false BOOL" json:"is_trade"`
	Token     string    `xorm:"not null VARCHAR(255)" json:"token"`
	Class     int64     `json:"type"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `xorm:"created" json:"created_at"`
	UpdatedAt time.Time `xorm:"updated" json:"updated_at"`
}

//region Remark:币种列表 Author:Beyond
func CoinList(page int, limit int, keywords string) *[]Coin {
	if keywords != "" {
		var coin = new([]Coin)
		err := databases.Orm.Table("coin").Where("name = ?", keywords).Limit(limit, page*limit).Find(coin)
		if err != nil {
			fmt.Println(err.Error())
		}
		return coin
	} else {
		var coin = new([]Coin)
		err := databases.Orm.Table("coin").Limit(limit, page*limit).Find(coin)
		if err != nil {
			fmt.Println(err.Error())
		}
		return coin
	}
}

//endregion

//region Remark:获取币种数据 Author:Beyond
func GetCoinInfo(coin_id int64) *Coin {
	coin := new(Coin)
	has, err := databases.Orm.Where("id= ?", coin_id).Get(coin)
	if err != nil {
		fmt.Println(err.Error())
	}
	if has == false {
		return nil
	}
	return coin
}

//endregion

//region Remark:获取币种下拉 Author:Beyond
func GetCoinList() *[]Coin {
	coin := new([]Coin)
	err := databases.Orm.Where("id > ?", 0).Select("id,code").Find(coin)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return coin
}

//endregion

//region Remark:根据code获取币种 Author:fang
func GetCoinByCode(code string) *Coin {
	coin := new(Coin)
	has, err := databases.Orm.Where("code = ?", code).Get(coin)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	if has == false {
		return nil
	}
	return coin
}

//endregion

//region Remark: 获取所有货币 Author; chijian
func GetCoinMap() map[int64]Coin {
	data := new([]Coin)
	err := databases.Orm.Find(data)
	if err != nil {
		fmt.Println(err.Error())
	}
	list := make(map[int64]Coin)
	for _, value := range *data {
		list[value.Id] = value
	}
	return list
}

func GetCoins() *[]Coin {
	data := new([]Coin)
	err := databases.Orm.Find(data)
	if err != nil {
		fmt.Println(err.Error())
	}
	return data
}

//endregion
