package main

import (
	"fmt"
	"reflect"
)

func main() {
	var str = "hello"
	GetType(str)
}

func GetType(e interface{}) {
	t := reflect.TypeOf(e)
	fmt.Println(t)

	val := reflect.ValueOf(e)
	fmt.Println(val)

	kind := val.Kind()
	fmt.Println(kind)
}
