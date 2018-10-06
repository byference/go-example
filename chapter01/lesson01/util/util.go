package util

import "fmt"

func Println(message ...interface{}) {
	fmt.Println(message)
}

func Add(i1, i2 int) int {
	return i1 + i2
}
