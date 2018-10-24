package main

import (
	"fmt"
	"time"
)

// Go 程序使用 error 值来表示错误状态。
// 自定义error
type MyError struct {
	When time.Time
	What string
}

// 类似java的toString方法
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
