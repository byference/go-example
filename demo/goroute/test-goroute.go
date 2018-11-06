package main

import (
	"fmt"
	"time"
)

func test_println(a int) {
	if a == 0 {
		fmt.Println("----------------")
		fmt.Println("第", a, "次.")
	}
	fmt.Println("第", a, "次.")
}

func main() {
	for i := 0; i < 10; i++ {
		go test_println(i)
		go test_println(i)
	}
	time.Sleep(time.Second)
}
