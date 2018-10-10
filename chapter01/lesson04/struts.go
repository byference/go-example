package main

import "fmt"

// 一个结构体struct就是一个字段的集合。
type MyIntList struct {
	X, Y int
}

var (
	p  = &MyIntList{1, 2} // 类型为 *Vertex
	v1 = MyIntList{1, 2}  // 类型为 Vertex
	v2 = MyIntList{X: 1}  // Y:0 被省略
	v3 = MyIntList{}      // X:0 和 Y:0
)

func main() {
	test2()
}

func test1() {
	list := MyIntList{1, 2}
	fmt.Println(list)
	fmt.Println("MyIntList中X的值：", list.X)
	// 结构体字段可以通过结构体指针来访问
	fmt.Println("MyIntList中X的指针：", &list.X)
}

func test2() {
	fmt.Println(p, v1, v2, v3)
}
