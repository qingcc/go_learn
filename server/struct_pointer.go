package main

import "fmt"

type Animal interface {
	Bark()
}
type Dog struct {
}

func (d Dog) Bark() {
	fmt.Println("dog")
}

type Cat struct {
}

func (c *Cat) Bark() {
	fmt.Println("cat")
}

func Bark(a Animal) {
	a.Bark()
}
func getDog() Dog {
	return Dog{}
}
func getCat() Cat {
	return Cat{}
}

func main() {

	dp := &Dog{}
	d := Dog{}
	dp.Bark() // (1) 通过
	d.Bark()  // (2) 通过
	Bark(dp)
	// (3) 通过，上面说了类型*Dog的方法集合包含接收者为*Dog和Dog的方法
	Bark(d) // (4) 通过

	cp := &Cat{}
	c := Cat{}
	cp.Bark() // (5) 通过
	c.Bark()  // (6) 通过
	Bark(cp)  // (7) 通过
	Bark(c)
	// (8) 编译错误，值类型Cat的方法集合只包含接收者为Cat的方法
	// 所以T并没有实现Animal接口

	getDog().Bark() // (9) 通过
	getCat().Bark()
	// (10) 编译错误，
	// 上面说了，getCat()是不可地址的
	// 所以不能调用接收者为*Cat的方法
}
