package databases

import (
	"fmt"
	"github.com/go-ini/ini"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	//"github.com/polaris1119/logger"
	"runtime"
	"strconv"
)

var Orm *xorm.Engine

func Init() error {
	conf, err := ini.Load("config/db.ini")
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Println(file, line) //***写入日志
	}
	key, _ := conf.Section("db").GetKey("host")
	host := key.Value()
	key, _ = conf.Section("db").GetKey("port")
	port := key.Value()
	key, _ = conf.Section("db").GetKey("user")
	user := key.Value()
	key, _ = conf.Section("db").GetKey("password")
	password := key.Value()
	key, _ = conf.Section("db").GetKey("dbname")
	dbname := key.Value()

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname) //pgsql 写法
	//psqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", user, password, host, port, dbname) //mysql 写法

	Orm, err = xorm.NewEngine("postgres", psqlInfo)
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Println(file, line) //***写入日志
		fmt.Println("连接数据库出错")  //***写入日志
		//logger.Errorln("连接pgsql出错")
	}

	//连接测试
	if err := Orm.Ping(); err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Println(file+":"+strconv.Itoa(line), "PostgresSql 数据库链接测试: %s", err.Error()) //***写入日志
	}

	//设置连接池的空闲数大小
	Orm.SetMaxIdleConns(1000000)
	//设置最大打开连接数
	Orm.SetMaxOpenConns(5000000)
	//名称映射规则主要负责结构体名称到表名和结构体field到表字段的名称映射
	Orm.SetTableMapper(core.SnakeMapper{})

	//连接成功
	_, file, line, _ := runtime.Caller(1)
	fmt.Println(file, line) //***写入日志

	return nil
}
