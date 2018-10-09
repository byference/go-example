package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	test4()

}

func os() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("linux.")
	default:
		fmt.Println(os)
	}
}

func test4() {
	t := time.Now()
	// 没有条件的switch同switch true一样
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}
}
