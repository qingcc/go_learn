package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	functions()
}

//os包中的一些常用的函数
func functions() {
	fmt.Println(os.Hostname())

	fmt.Println(os.Getwd())

	fmt.Println(os.Getuid())

	fmt.Println(os.Geteuid())
	fmt.Println(os.Getgid())
	fmt.Println(os.Getegid())
	fmt.Println(os.Getpid())
	fmt.Println(os.Getppid())

	fmt.Println("------")
	fmt.Println(os.Getenv("GOPATH"))

	fmt.Println(strings.Join(os.Environ(), "\r\n"))

	//把字符串中带${var}或$var替换成指定指符串
	fmt.Println(os.Expand("${1} ${2} ${3}", func(k string) string {
		mapp := map[string]string{
			"1": "111",
			"2": "222",
			"3": "333",
		}
		return mapp[k]
	}))
}
