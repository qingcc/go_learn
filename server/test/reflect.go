package main

import (
	"fmt"
	"reflect"
)

// patch函数把diff apply到input上。
// input必须是一个指针。
// diff的key为input的field名称, value为更新后的值
func patch(input interface{}, diff map[string]interface{}) {
	entityValue := reflect.Indirect(reflect.ValueOf(input))
	fmt.Println("va:", entityValue)
	for field, value := range diff {
		fieldValue := entityValue.FieldByName(field)

		fmt.Println("field:", field, "val:", fieldValue)
		if value == nil {
			// diff中value为nil,将fielValue设为默认值
			fieldValue.Set(reflect.Zero(fieldValue.Type()))
		} else {
			fieldValue.Set(reflect.ValueOf(value))
		}
	}
}

type testStruct struct {
	A int     `xorm:"-" json:"a"`
	B string  `json:"b"`
	C float64 `json:"c"`
}

func main() {
	//input := &testStruct{A: 10, B: "hello", C: 3.14}
	//patch(input, map[string]interface{}{"A": 11, "B": "world", "C": nil})
	//fmt.Printf("%+v", input) // &{A:11 B:world C:0}, A被更新到11， B被更新到world，C被更新到float64的default值
	reflect_main()
}

type Student struct {
	Name    string
	Age     int
	Sex     uint // 0-女性，1-男性
	Address string
}

func (stu Student) Print() {
	sex := "女"
	if stu.Sex == 1 {
		sex = "男"
	}
	fmt.Printf("姓名：%s，年龄：%d，性别：%s，地址：%s", stu.Name, stu.Age, sex, stu.Address)
}

func (stu Student) Say(content string) string {
	return fmt.Sprintf("%s说：%s", stu.Name, content)
}

func reflect_main() {
	// ============================================================操作基本数据类型
	var a1 int = 1
	var a2 string = "中国"
	var a3 = [...]byte{1, 2, 3}                    // 数组
	var a4 = []int{5, 6, 7}                        // 切片
	var a5 = map[string]int{"china": 5, "usa": 40} // Map

	// 转化成reflect对象，reflect.Type和reflect.Value
	// TypeOf()返回的数据类型是*reflect.rtype，ValueOf()返回的数据类型是reflect.Value
	t1 := reflect.TypeOf(a1)
	v1 := reflect.ValueOf(a1)
	t2 := reflect.TypeOf(a2)
	v2 := reflect.ValueOf(a2)
	t3 := reflect.TypeOf(a3)
	v3 := reflect.ValueOf(a3)
	t4 := reflect.TypeOf(a3)
	v4 := reflect.ValueOf(a4)
	t5 := reflect.TypeOf(a3)
	v5 := reflect.ValueOf(a5)
	fmt.Println("a1====", t1, v1)
	fmt.Println("a2====", t2, v2)
	fmt.Println("a3====", t3, v3)
	fmt.Println("a4====", t4, v4)
	fmt.Println("a5====", t5, v5)
	fmt.Println()

	// 取值
	fmt.Println("a1的值：", v1.Int())
	fmt.Println("a2的值：", v2.String())
	fmt.Println("a3中下标为1的元素的值：", v3.Index(1))
	fmt.Println("a4中下标为1的元素的值：", v4.Index(1))
	fmt.Println("a4中取[1,3)的子切片：", v4.Slice(1, 3))
	fmt.Println("a5中所有key：", v5.MapKeys())
	fmt.Print("遍历a5中的value：")
	fmt.Println("#########################################################################################################")
	for i, key := range v5.MapKeys() {
		fmt.Print(key, "====", v5.MapIndex(key), "\t")
		fmt.Println("i:", i, "key:", key.String())
	}
	fmt.Println()
	fmt.Println()

	// 获取类型
	fmt.Println("a1的类型：", v1.Type())
	fmt.Println("a2的类型：", v2.Type())
	fmt.Println("a3的类型：", v3.Type())
	fmt.Println("a4的类型：", v4.Type())
	fmt.Println("a5的类型：", v5.Type())
	fmt.Println()

	// Kind类型判断，Kind()
	fmt.Println("a1的类型是否为int：", v1.Kind() == reflect.Int)
	fmt.Println("a2的类型是否为string：", v2.Kind() == reflect.String)
	fmt.Println("a3的类型是否为array：", v3.Kind() == reflect.Array)
	fmt.Println("a4的类型是否为slice：", v4.Kind() == reflect.Slice)
	fmt.Println("a5的类型是否为map：", v5.Kind() == reflect.Map)
	fmt.Println()

	// 接口类型变量，Interface是获取该value的值,返回的是一个interface对象
	fmt.Printf("a1====%T\t%v\n", v1.Interface(), v1.Interface())
	fmt.Printf("a2====%T\t%v\n", v2.Interface(), v2.Interface())
	fmt.Printf("a3====%T\t%v\n", v3.Interface(), v3.Interface())
	fmt.Printf("a4====%T\t%v\n", v4.Interface(), v4.Interface())
	fmt.Printf("a5====%T\t%v\n", v5.Interface(), v5.Interface())
	fmt.Println()

	// 判断是否可以修改
	fmt.Println("a1通过反射是否可以进行修改：", v1.CanSet())
	fmt.Println("a2通过反射是否可以进行修改：", v2.CanSet())
	fmt.Println("a3通过反射是否可以进行修改：", v3.CanSet())
	fmt.Println("a4通过反射是否可以进行修改：", v4.CanSet())
	fmt.Println("a5通过反射是否可以进行修改：", v5.CanSet())
	fmt.Println()

	// 修改，必须传指针，且调用Elem()
	vv1 := reflect.ValueOf(&a1).Elem()
	fmt.Printf("vv1的类型：%T，a1现在是否可以通过：%v\n", vv1, vv1.CanSet())
	vv2 := reflect.ValueOf(&a2).Elem()
	vv3 := reflect.ValueOf(&a3).Elem()
	vv4 := reflect.ValueOf(&a4).Elem()
	vv5 := reflect.ValueOf(&a5).Elem()

	vv1.SetInt(100) // 修改
	vv2.SetString("美国")
	var temp byte = 90
	vv3.Index(0).Set(reflect.ValueOf(temp)) // 必须传reflect.Value类型
	vv4.Index(0).SetInt(900)
	// 必须传reflect.Value类型
	vv5.SetMapIndex(reflect.ValueOf("china"), reflect.ValueOf(111))

	fmt.Println("a1修改后的值：", a1)
	fmt.Println("a2修改后的值：", a2)
	fmt.Println("a3修改后的值：", a3)
	fmt.Println("a4修改后的值：", a4)
	fmt.Println("a5修改后的值：", a5)

	// ============================================================操作结构体
	stu := Student{"李四", 18, 1, "中国北京市天安门10000号"}

	// 转化成reflect对象，reflect.Type和reflect.Value
	st1 := reflect.TypeOf(stu)
	sv1 := reflect.ValueOf(stu)
	fmt.Println(st1, "====", sv1)

	// 获取结构体名称
	fmt.Println(st1.Name())

	// 判断Kind类型
	fmt.Println(st1.Kind() == reflect.Struct)

	// 获取结构体中字段数量
	fmt.Println(st1.NumField())

	// 获取结构体中每个字段的值
	for i := 0; i < st1.NumField(); i++ {
		fieldName := st1.Field(i).Name // 取字段名
		fieldType := st1.Field(i).Type
		fieldVal := sv1.Field(i).Interface() // 取值
		fmt.Printf("字段名：%v，类型：%v，值：%v\n", fieldName, fieldType, fieldVal)
	}

	// 获取结构体的方法数量
	fmt.Println(st1.NumMethod())
	// 遍历所有方法的名称和类型
	for i := 0; i < st1.NumMethod(); i++ {
		method := st1.Method(i)
		fmt.Println(method.Name, "====", method.Type)
	}

	// 通过反射执行方法
	m1 := sv1.MethodByName("Print")
	m1.Call(nil) // 不带参数和返回值
	fmt.Println()

	m2 := sv1.MethodByName("Say")
	params := []reflect.Value{reflect.ValueOf("你好啊！")}
	res := m2.Call(params) // 带参数和返回值，参数是reflect.Value的切片，返回值也一样
	fmt.Println(res[0].String())
}
