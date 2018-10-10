package main

import "fmt"

func main() {

	// 函数也是值
	fun := func(str string) string {
		return "new " + str
	}
	fmt.Println(fun("joy"))
}
