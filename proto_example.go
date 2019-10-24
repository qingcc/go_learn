/*
*protobuf using example
*
 */

package main

import (
	"example"
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"os"
)

func person() {
	// 为 AllPerson 填充数据
	p1 := example.Person{
		Id:   *proto.Int32(1),
		Name: *proto.String("xieyanke"),
	}

	p2 := example.Person{
		Id:   2,
		Name: "gopher",
	}

	all_p := example.AllPerson{
		Per: []*example.Person{&p1, &p2},
	}

	// 对数据进行序列化
	data, err := proto.Marshal(&all_p)
	if err != nil {
		fmt.Print("Mashal data error:", err)
	}

	// 对已经序列化的数据进行反序列化
	var target example.AllPerson
	err = proto.Unmarshal(data, &target)
	if err != nil {
		fmt.Print("UnMashal data error:", err)
	}

	println(target.Per[0].Name) // 打印第一个 person Name 的值进行反序列化验证
}

func write() {
	p1 := &example.Personal{
		Id:   1,
		Name: "小张",
		Phones: []*example.Phone{
			{Type: example.PhoneType_HOME, Number: "111111111"},
			{Type: example.PhoneType_WORK, Number: "222222222"},
		},
	}
	p2 := &example.Personal{
		Id:   2,
		Name: "小王",
		Phones: []*example.Phone{
			{Type: example.PhoneType_HOME, Number: "333333333"},
			{Type: example.PhoneType_WORK, Number: "444444444"},
		},
	}

	//创建地址簿
	book := &example.ContactBook{}
	book.Persons = append(book.Persons, p1)
	book.Persons = append(book.Persons, p2)

	//编码数据
	data, _ := proto.Marshal(book)
	//把数据写入文件
	ioutil.WriteFile("./test.txt", data, os.ModePerm)
}

func read() {
	//读取文件数据
	data, _ := ioutil.ReadFile("./test.txt")
	book := &example.ContactBook{}
	//解码数据
	proto.Unmarshal(data, book)
	for _, v := range book.Persons {
		fmt.Println(v.Id, v.Name)
		for _, vv := range v.Phones {
			fmt.Println(vv.Type, vv.Number)
		}
	}
}

func main() {
	person()
	write()
	read()
}
