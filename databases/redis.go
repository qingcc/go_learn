package databases

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/go-ini/ini"
	"strconv"
	"time"
)

var (
	host      string
	port      string
	password  string
	dbname    int
	maxidle   int
	maxactive int
)

//
func RedisClient() *redis.Pool {
	//读取配置文件
	config, err := ini.Load("config/redis.ini")
	if err != nil {
		fmt.Println(err.Error())
	}
	key, _ := config.Section("session").GetKey("host")
	host = key.Value()
	key, _ = config.Section("session").GetKey("port")
	port = key.Value()
	key, _ = config.Section("session").GetKey("password")
	password = key.Value()
	key, _ = config.Section("session").GetKey("dbname")
	dbname, _ = strconv.Atoi(key.Value())
	key, _ = config.Section("session").GetKey("maxidle")
	maxidle, _ = strconv.Atoi(key.Value())
	key, _ = config.Section("session").GetKey("maxactive")
	maxactive, _ = strconv.Atoi(key.Value())
	//连接redis
	RedisClient := &redis.Pool{
		//从配置文件获取 maxidle 和 maxactive 取不到则使用默认
		MaxIdle:     maxidle,
		MaxActive:   maxactive,
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", host+":"+port)
			if err != nil {
				fmt.Println(err.Error())
				return nil, err
			}
			if _, err := c.Do("AUTH", password); err != nil {
				c.Close()
				fmt.Println(err)
				return nil, err
			}

			//选择db
			c.Do("SELECT", dbname)
			return c, nil
		},
	}
	return RedisClient
}
