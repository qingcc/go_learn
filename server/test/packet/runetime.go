package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	exit := make(chan int)
	go func() {
		defer close(exit)
		go func() {
			fmt.Println("b")
		}()
	}()

	for i := 0; i < 10; i++ {
		fmt.Println("a:", i)

		if i == 4 {
			runtime.Gosched() //切换任务
		}
	}
	<-exit

}

//package main
//
//import (
//	"fmt"
//	"runtime"
//)
//
//func showNumber(i int) {
//	fmt.Println(i)
//}
//
//func main() {
//	runtime.GOMAXPROCS(1)
//	for i := 0; i < 10; i++ {
//		go showNumber(i)
//	}
//
//	runtime.Gosched()
//	runtime.GC()
//	fmt.Println("Haha")
//}

//
//package main
//
//import (
//	"fmt"
//	"runtime"
//)
//
//func say(s string) {
//	for i := 0; i < 2; i++ {
//
//		runtime.Gosched()
//		fmt.Println(s)
//
//	}
//}
//
//func main() {
//	//runtime.GOMAXPROCS(1)
//	//go say("world")
//	//say("hello")
//
//	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
//	for _, value := range data {
//		fmt.Println(value)
//	}
//}
