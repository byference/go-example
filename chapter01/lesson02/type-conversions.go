package lesson02

import "fmt"

func main() {
	var x, y int = 3, 4
	// 数值转换
	var f1 float64 = float64(x)
	// 数值转换简写
	f2 := float64(y)
	fmt.Println(f1, f2)
}
