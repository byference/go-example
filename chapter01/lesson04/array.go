package main

import "fmt"

func main() {

	test8()

}

func test3() {
	// 类型 [n]T 是一个有 n 个类型为 T 的值的数组
	// 定义变量 a 是一个有2个string的数组
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
}

func test4() {
	ints := []int{1, 2, 3, 4}
	fmt.Println("数组ints的长度为：", len(ints))
}

// 对 slice 切片
// 表示从 lo 到 hi-1 的 slice 元素，含两端
func test5() {
	// ints := []int{1, 2, 3, 4}
	ints := [5]int{}
	fmt.Println(ints[0:3])

}

func test6() {

	// b := make([]int, 5) // len(b)=5, cap(b)=5
	b := make([]int, 0, 5) // len(b)=0, cap(b)=5
	fmt.Println(b)
	b = b[:cap(b)] // len(b)=5, cap(b)=5
	fmt.Println(b)
	b = b[1:] // len(b)=4, cap(b)=4
	fmt.Println(b)
}

// slice 的零值是 `nil`
func test7() {
	var z []int
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil!")
	}
}

func test8() {
	var a []int
	printSlice("a", a)

	// append works on nil slices.
	a = append(a, 0)
	printSlice("a", a)

	a = append(a, 1)
	printSlice("a", a)

}

func printSlice(s string, x []int) {

	fmt.Printf("%s: len=%d, cap=%d v=%v\n", s, len(x), cap(x), x)
}
