package main

import (
	"blog_go/database/sql"
	"fmt"
)

func main() {
	test()

}

func test() {
	db, err := sql.Open("odbc", "DSN=test;")
	defer db.Close()

	checkErr(err)

	stmt, err := db.Prepare("select name from table")
	defer stmt.Close()
	checkErr(err)
	rows, err := stmt.Query()
	defer rows.Close()
	checkErr(err)
	for rows.Next() {
		var name string

		_ = rows.Scan(&name)
		fmt.Println(name)
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}
