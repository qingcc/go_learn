package model

import (
	"github.com/qingcc/goblog/databases"
	"fmt"
	"github.com/go-xorm/xorm"
	"time"
	"github.com/qingcc/goblog/util"
)

type CoinPrice1Min struct {
	Id        int64     `xorm:"pk autoincr BIGINT"  json:"id"`
	CoinId    int64     `xorm:"not null int" json:"coin_id"`        //货币id
	High      float64   `xorm:"not null decimal(12)"`               //最高价
	Low       float64   `xorm:"not null decimal(12)"`               //最低价
	Open      float64   `xorm:"not null decimal(12)"`               //开盘价
	Close     float64   `xorm:"not null decimal(12)"`               //收盘价
	Count     float64   `xorm:"not null decimal(18)"`               //交易次数
	Amount    float64   `xorm:"not null decimal(18)" json:"amount"` //以基础币种计量的交易量
	Vol       float64   `xorm:"not null decimal(12)" json:"vol"`    //以报价币种计量的交易量
	Time      time.Time `xorm:"not null timestamp"`
	Grow      float64   `xorm:"not null decimal(12)" json:"grow"` //价格增长额度
	Rise      float64   `xorm:"not null decimal(12)" json:"rise"` //价格涨幅
	UpdatedAt time.Time `xorm:"updated"`
}

var CoinPrice = map[int]K{
	1: {table: "coin_price1_min", interval_time: time.Minute, time_type: 1},
	2: {table: "coin_price15_min", interval_time: time.Minute * 15, time_type: 1},
	3: {table: "coin_price2_hour", interval_time: time.Hour * 2, time_type: 2},
	4: {table: "coin_price4_hour", interval_time: time.Hour * 4, time_type: 2},
	5: {table: "coin_price1_day", interval_time: time.Hour * 24, time_type: 3},
}

type K struct {
	table         string
	interval_time time.Duration
	time_type     int
}

//region Remark: 获取最新的1条数据 Author; chijian
func GetNewCoinPrice(coin_id int64, time_type int, limit int) *CoinPrice1Min {
	item := &CoinPrice1Min{}
	res, err := databases.Orm.Table(CoinPrice[time_type].table).Where("coin_id = ?", coin_id).Desc("time").Limit(limit, 0).Get(item)
	if err != nil {
		fmt.Println(err.Error())
	}
	if res == false {
		return nil
	}
	return item
}

//endregion

//region Remark: 获取指定时间内的数据 Author; chijian
func GetCoinPriceData(coin_id int64, time_type int) *[]CoinPrice1Min {
	data := new([]CoinPrice1Min)
	err := databases.Orm.Table(CoinPrice[time_type].table).Where("coin_id = ?", coin_id).Find(data)
	if err != nil {
		fmt.Println(err.Error())
	}
	return data
}

//endregion

//region Remark: 更新k线数据 Author; chijian
func UpdateCoinPriceData(db *xorm.Session, coin_id int64, coin_price *CoinPrice1Min) bool {
	res := GetRiseAndGrow(db, coin_id, coin_price) //获取价格增长和价格涨幅
	if res == false {
		return false
	}

	for i := 1; i < 6; i++ {
		res := updateData(db, coin_id, i, coin_price) //更新
		if res == false {
			return false
		}
	}
	return true
}

//获取价格增长和价格涨幅
func GetRiseAndGrow(db *xorm.Session, coin_id int64, coin_price *CoinPrice1Min) bool {
	t1 := time.Now()
	t1 = time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, time.Local)
	item := new(CoinPrice1Min)
	res, err := db.Table(CoinPrice[5].table).Where("coin_id = ? and time < ?", coin_id, t1).Desc("time").Get(item)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	if res == false { //该货币昨天没有数据
		res, err := db.Table(CoinPrice[1].table).Where("coin_id = ? and time > ?", coin_id, t1).Asc("time").Get(item)
		if err != nil {
			fmt.Println(err.Error())
			return false
		}
		if res == false { //该货币没有数据
			coin_price.Grow = 0
			coin_price.Rise = 0
		} else { //该货币在当天有数据
			coin_price.Grow = util.RoundFloat(coin_price.Close-item.Open, 4)
			coin_price.Rise = util.RoundFloat(coin_price.Grow/item.Open, 4)
		}
	} else { //该货币前一天有数据
		coin_price.Grow = util.RoundFloat(coin_price.Close-item.Close, 4)
		coin_price.Rise = util.RoundFloat(coin_price.Grow/item.Close, 4)
	}
	return true
}

//更新
func updateData(db *xorm.Session, coin_id int64, time_type int, coin_price *CoinPrice1Min) bool {
	data := &CoinPrice1Min{}
	table := CoinPrice[time_type].table

	res, err := db.Table(table).Where("coin_id = ? and time > ?", coin_id, time.Now().Add(-CoinPrice[time_type].interval_time)).Desc("time").Get(data)
	if err != nil {
		fmt.Println(err.Error())
	}

	var has int64
	if res { //有数据,间隔时间内
		has, err = db.Table(table).Id(data.Id).Cols("close", "high", "low", "grow", "rise", "count", "amount", "vol").Update(coin_price)
	} else { //超出时间,新增1条数据
		coin_price.Id = 0
		coin_price.Time = util.GetTimeType(time.Now(), CoinPrice[time_type].time_type)
		has, err = db.Table(table).Insert(coin_price)

		//更新货币价格coin表,每15min更新1次
		if time_type == 2 {
			//更改货币最新价格
			coin := &Coin{Price: coin_price.Open}
			has, err := db.Where("id = ?", coin_id).Cols("price").Update(coin)
			if err != nil {
				fmt.Println(err.Error())
				db.Rollback()
				return false
			}
			if has < 1 {
				db.Rollback()
				return false
			}

		}
	}
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	if has < 1 {
		return false
	}
	return true
}

//endregion
