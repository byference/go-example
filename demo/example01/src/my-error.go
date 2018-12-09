package main

import (
	"fmt"
	"time"
)

type MyError struct {
	Timestamp time.Time
	Msg       string
}

func (err *MyError) Error() string {
	return fmt.Sprintf("Timestamp:%v, Msg:%v", err.Timestamp, err.Msg)
}

func main() {
	err := TestError()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("success")
}

func TestError() error {
	return &MyError{
		Timestamp: time.Now().Local(),
		Msg:       "something is error!",
	}
}
