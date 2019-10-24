package main

import "sync"

/*
* todo 加锁 Mutex 和 channel 性能对比
* go test -bench=.
 */

var mutex = sync.Mutex{}
var ch = make(chan bool, 1)

func UseMutex() {
	mutex.Lock()
	mutex.Unlock()
}
func UseChan() {
	ch <- true
	<-ch
}
