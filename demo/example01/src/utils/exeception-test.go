package utils

import (
	"errors"
	"fmt"
)

func MyError() error {
	return errors.New("MyError happened")
}

func MyErrorTest() {

	err := MyError()
	if err != nil {
		panic(err)
	}
}

func Error() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("*********测试异常: ", err)
		}
	}()

	a := 0
	b := 100 / a
	fmt.Println(b)
}
