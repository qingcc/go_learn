package main

import (
	"blog_go/databases"
	"fmt"
)

// bingone
func main() {

	toTruncateTable()
	fmt.Println("清空成功")
	return
}

func toTruncateTable() {
	rebuildSequence("chat")
}

func rebuildSequence(table string) {
	gsql := "delete from " + table + " where id > 0" //绑定序列
	exec(gsql)
}

func exec(sql string) {
	_, err := databases.Orm.Exec(sql)
	if err != nil {
		fmt.Println("sql 执行失败::", sql)
	}
}
