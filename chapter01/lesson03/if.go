package main

import "fmt"

func main() {

	fmt.Println(max2(3, 4, 5))
}

func max1(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 跟for一样，if语句可以在条件之前执行一个简单的语句，变量的作用域仅在if范围之内。
func max2(a, b, c int) int {
	if v := max1(a, b); v > c {
		return v
	} else {
		fmt.Println("this is v:", v)
	}
	return c
}
