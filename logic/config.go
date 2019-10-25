package logic

import (
	"github.com/qingcc/go_learn/databases"
	"github.com/garyburd/redigo/redis"
	"github.com/qingcc/go_learn/model"
	"github.com/qingcc/go_learn/util"
)

type ConfigLogic struct{}

var DefaultConfig = ConfigLogic{}

func (self ConfigLogic) ReadConfig(name string) string {
	//判断redis是否有缓存数据
	redis_key := "admin:config:" + name
	if res, _ := util.Exists(redis_key); res {
		value, _ := redis.String(util.Get(redis_key))
		return value
	}

	//查询
	item := new(model.Config)
	databases.Orm.Where("name = ?", name).Get(item)
	util.Set(redis_key, item.Value, 60*60)
	return item.Value
}

func (self ConfigLogic) SetConfig(data []model.Config) {
	for _, value := range data {
		item := new(model.Config)
		res, err := databases.Orm.Where("name = ?", value.Name).Get(item)
		util.CheckErr(err)

		if res { //存在, 更新
			databases.Orm.Cols("value").Where("name = ?", value.Name).Update(value)
		} else { //不存在, 插入
			databases.Orm.Insert(value)
		}
	}
	util.DelKeyByPrefix("admin:config:")
}
