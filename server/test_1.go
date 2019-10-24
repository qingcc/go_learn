package main

import (
	"fmt"
	"time"
)

func main() {
	//item := map[string]interface{}{
	//	"id":    33123,
	//	"token": "askdflsalkfd",
	//}
	//_item, err := json.Marshal(item)
	//fmt.Println("err:", err, "item:", string(_item))
	//
	//str := ","
	//b := strings.Contains(str, ",32,")
	//fmt.Println("b:", b)
	//fmt.Println("_len:", len(strings.Split(str, ",")))
	////fmt.Println("len:", len(strings.Split(str[1:len(str)-1], ",")))
	//fmt.Println("str:", str, "len:", len(strings.Split(",32,", ",")))
	//test()

	t := time.Now().Add(1e9 * 60)
	fmt.Println("time:", time.Now(), "\ntime:", t)
	a := 2343444444444.2333333
	fmt.Println(fmt.Sprintf("%.2f", a))
}

func test() {

	d := time.Duration(time.Second * 1)

	t := time.NewTimer(d) //设置定时1秒种
	defer t.Stop()
	i := 1
	for {
		<-t.C

		fmt.Println("timeout...", i)
		i = i + 1
		// need reset
		t.Reset(time.Second * 1) //重置定时器
	}
}
func test1() {
	//初始化断续器,间隔2s
	var ticker *time.Ticker = time.NewTicker(100 * time.Millisecond)

	//num为指定的执行次数
	num := 2
	c := make(chan int, num)
	go func() {
		for t := range ticker.C {
			c <- 1
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(time.Millisecond * 1500)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
