package main

// 省略了循环条件，循环就不会结束，因此可以用更简洁地形式表达死循环

import "fmt"

func main() {

	test3()

}

// for循环
func test1() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum++
	}
	fmt.Println(sum)
}

// 允许前置、后置语句为空
func test2() {
	sum := 0
	for sum < 100 {
		sum++
	}
	fmt.Println(sum)
}

// 等同于java中的while
func test3() {
	sum := 0
	for sum < 100 {
		sum++
	}
	fmt.Println(sum)
}
