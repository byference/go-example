package lesson02

import "fmt"

const SUCCESS = 1

func main() {

	// 常量的定义与变量类似，只不过使用 const 关键字。
	//常量可以是字符、字符串、布尔或数字类型的值。
	//常量不能使用 := 语法定义。
	const big = 1 << 10
	const (
		str1 string = "hello,"
		str2 string = "world"
	)

	fmt.Println(big)

}
