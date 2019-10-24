package main

import (
	"fmt"

	simplejson "github.com/bitly/go-simplejson"
)

func main() {
	js, err := simplejson.NewJson([]byte(`{
    "test": {
      "string_array": ["asdf", "ghjk", "zxcv"],
      "array": [1, "2", 3],
      "arraywithsubs": [{"subkeyone": 1},
      {"subkeytwo": 2, "subkeythree": 3}],
      "int": 10,
      "float": 5.150,
      "bignum": 9223372036854775807,
      "string": "simplejson",
      "bool": true
    }
  }`))
	if err != nil {
		panic("json format error")
	}

	//获取某个字段值
	s, err := js.Get("test").Get("string").String()
	if err != nil {
		panic(err)
	}
	fmt.Println(s)

	//检查某个字段是否存在
	_, ok := js.Get("test").CheckGet("string2")
	if ok {
		fmt.Println("存在！")
	} else {
		fmt.Println("不存在")
	}

	//数组的操作
	arr, err := js.Get("test").Get("string_array").StringArray()
	if err != nil {
		panic(err)
	}
	for _, v := range arr {
		fmt.Printf("%s\n", v)
	}

	arr2, err := js.Get("test").Get("array").Array()
	if err != nil {
		panic(err)
	}

	for _, v := range arr2 {
		fmt.Printf("%T:%v\n", v, v)
	}

	arr3 := js.Get("test").Get("arraywithsubs").GetIndex(1).MustMap()
	// if err != nil {
	//  panic(err)
	// }
	fmt.Printf("%v", arr3)

}
