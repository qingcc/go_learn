package main

import (
	"blog_go/databases"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"blog_go/model"
	"net/http"
	"strings"
	"time"
	"blog_go/util"
)

func main() {
	//toLoadStartData()
	//time.Sleep(time.Hour)
	for {
		getPrice()
		//if !has_wrong {
		//	fmt.Println("写入成功")
		//}
		time.Sleep(time.Second * 1)
	}
}

func getPrice() {
	coins := model.GetCoins()
	for _, value := range *coins {
		if value.Code == "SUM" {
			continue
		}
		go changeData(strings.ToLower(value.Code), value.Id)
	}
}

var Url = "https://api.huobi.pro"
var has_wrong = false

//插入,更改数据
func changeData(coin_name string, coin_id int64) {
	url := Url + "/market/detail/merged?symbol=" + coin_name + "usdt"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("请求获取货币价格出错!\n", "time:", time.Now())
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取货币价格返回值出错!\n", "time:", time.Now())
		return
	}
	data := &response{}
	json.Unmarshal(body, &data)

	if data.Status == "ok" {
		item := new(model.CoinPrice1Min)
		item.CoinId = coin_id
		item.Open = data.Tick.Open
		item.Close = data.Tick.Close
		item.High = data.Tick.High
		item.Low = data.Tick.Low
		item.Count = data.Tick.Count
		item.Amount = data.Tick.Amount
		item.Vol = data.Tick.Vol
		if insert(data, coin_id, item) == false {
			fmt.Println("写入失败")
			has_wrong = true
		}
	} else {
		fmt.Println("无该币种:", coin_name, "time:", time.Now())
	}
}

func insert(data *response, coin_id int64, item *model.CoinPrice1Min) bool {
	//开启事物
	db := databases.Orm.NewSession()
	err := db.Begin()
	if err != nil {
		fmt.Println(err.Error())
	}

	res := model.UpdateCoinPriceData(db, coin_id, item)
	if res == false {
		db.Rollback()
		return false
	}

	db.Commit()
	return true
}

type Data struct {
	Id     int64
	Ask    []float64
	Bid    []float64
	High   float64 //最高价
	Low    float64 //最低价
	Open   float64 //开盘价
	Close  float64 //收盘价
	Count  float64 //交易次数
	Amount float64 //以基础币种计量的交易量
	Vol    float64 //以报价币种计量的交易量
}

type response struct {
	Status string
	Ch     string
	Ts     int
	Tick   Data
}

//region Remark: 初始化数据 Author; chijian
func toTruncateTable() {
	util.TruncateTable("coin_price1_min", "")
	util.TruncateTable("coin_price15_min", "")
	util.TruncateTable("coin_price2_hour", "")
	util.TruncateTable("coin_price4_hour", "")
	util.TruncateTable("coin_price1_day", "")
}

//endregion
