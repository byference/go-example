package main

import "fmt"

var list = []int{1, 2, 3, 4, 5, 6}

func main() {

	test10()
}

// for 循环的 range 格式可以对 slice 或者 map 进行迭代循环
func test9() {
	for i, v := range list {
		fmt.Printf("index%d - %d\n", i, v)
	}
}

// 可以通过赋值给 _ 来忽略序号和值
func test10() {
	for _, v := range list {
		fmt.Printf("vaule= %d\n", v)
	}
}
