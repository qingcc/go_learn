package main

import (
	"fmt"
	"sync"
)

/*
* TODO  流水线pipeline的定义：
* 1.通过 inbound channels 从上游接收数据
* 2.对接收到的数据执行一些操作，通常会生成新的数据
* 3.将新生成的数据通过 outbound channels 发送给下游

× TODO 期望结果：当后一个阶段不需要数据时，上游阶段能够停止生产。
*
* todo 实现： 使用一个全局 channel done， 当下游关闭该channel done时， 上游的所有goroutine全部结束
× todo 原理： 在一个已关闭 channel 上执行接收操作(<-ch)总是能够立即返回，返回值是对应类型的零值
×
*/

func main() {
	// 设置一个 全局共享的 done channel，
	// 当流水线退出时，关闭 done channel
	// 所有 goroutine接收到 done 的信号后，
	// 都会正常退出。
	done := make(chan struct{})
	defer close(done)

	in := gen(done, 2, 3, 4, 5, 6, 7, 8, 9, 10) //gen方法，将多个输入写入一个 channel

	// 将 sq 的工作分发给两个goroutine
	// 这两个 goroutine 均从 in 读取数据
	c1 := sq(done, in) //sq方法，将扇入的值扇出，并计算平方数， 写入一个channel
	c2 := sq(done, in)

	// 消费 outtput 生产的第一个值
	out := merge(done, c1, c2) //merge方法，将多个channel里面的所有扇入的数据都扇出并写入一个channel中
	fmt.Println(<-out)         // 4 or 9
	fmt.Println(<-out)         // 4 or 9

	// defer 调用时，done channel 会被关闭。
}

func gen(done <-chan struct{}, nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			//out <- n
			//优化
			select {
			case out <- n:
			case <-done: // TODO 在一个已关闭 channel 上执行接收操作(<-ch)总是能够立即返回，返回值是对应类型的零值
			}
		}
		close(out)
	}()

	return out
}

func sq(done chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			//out <- n * n
			//优化
			select {
			case out <- n * n:
			case <-done: // TODO 在一个已关闭 channel 上执行接收操作(<-ch)总是能够立即返回，返回值是对应类型的零值
			}
		}
		close(out)
	}()
	return out
}

func merge(done <-chan struct{}, cs ...<-chan int) <-chan int {
	out := make(chan int)
	wg := sync.WaitGroup{}

	output := func(c <-chan int) {
		for n := range c {
			//out <- n
			//优化
			select {
			case out <- n:
			case <-done: // TODO 在一个已关闭 channel 上执行接收操作(<-ch)总是能够立即返回，返回值是对应类型的零值
			}
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
