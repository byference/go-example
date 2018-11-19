package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 互斥锁
var rwLock = sync.RWMutex{}

var count int32 = 0

func Test() {
	atomic.AddInt32(&count, 1)
}

func main() {
	for i := 0; i < 10000; i++ {
		rwLock.Lock()
		go Test()
		rwLock.Unlock()
	}

	time.Sleep(time.Second)

	rwLock.Lock()
	fmt.Println("count: ", atomic.LoadInt32(&count))
	rwLock.Unlock()

}
