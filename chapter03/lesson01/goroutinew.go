package lesson01

import (
	"fmt"
	"time"
)

// goroutine 是由 Go 运行时环境管理的轻量级线程
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

// goroutine 在相同的地址空间中运行，因此访问共享内存必须进行同步。
// sync 提供了这种可能，不过在 Go 中并不经常用到
func main() {
	// 开启一个新的 goroutine 执行
	go say("world")
	say("hello")
}
