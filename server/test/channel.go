//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//func producer(nums ...int) <-chan int {
//	out := make(chan int)
//	go func() {
//		defer close(out)
//		for _, n := range nums {
//			out <- n
//		}
//	}()
//	return out
//}
//
//func square(inCh <-chan int) <-chan int {
//	out := make(chan int)
//	go func() {
//		defer close(out)
//		for n := range inCh {
//			out <- n * n
//		}
//	}()
//
//	return out
//
//}
//func main() {
//	in := producer(1, 2, 3, 4)
//	fmt.Println("111", in)
//	time.Sleep(time.Second * 5)
//	ch := square(in)
//	fmt.Println(<-ch, "---")
//	time.Sleep(time.Second * 2)
//	// consumer
//	for ret := range ch {
//		fmt.Printf("%3d", ret)
//	}
//	fmt.Println()
//}

package main

import (
	"fmt"
)

func main() {
	mapResults := make(map[int]string)
	var arrResults [][]string

	count := 5
	for i := 0; i < count; i++ {
		valueStr := fmt.Sprintf("this is %d", i)
		mapResults[i] = valueStr
		var tmpArr []string

		for j := 0; j < 15; j++ {
			tmpArr = append(tmpArr, "a")

		}

		arrResults = append(arrResults, tmpArr)

	}

	fmt.Println(mapResults)
	fmt.Println(arrResults)
}
