package main

// Go 没有类。然而仍然可以在结构体类型上定义方法。
import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

type MyInt int

// 方法接收者出现在func关键字和方法名之间的参数中。
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// v是拷贝值
func (v Vertex) Add() float64 {
	return v.X + v.Y
}

// 你可以对包中的 任意 类型定义任意方法，而不仅仅是针对结构体。
//但是，不能对来自其他包的类型或基础类型定义方法。
func (i MyInt) testMyInt() int {
	return int(i)
}

// v是指针
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	test2()
}

func test1() {
	param := Vertex{3, 4}
	fmt.Println(param)
	v := &param
	fmt.Println(v.Abs())
	fmt.Println("Add->", (&Vertex{3, 4}).Add())
}

/*
方法可以与命名类型或命名类型的指针关联。
有两个原因需要使用指针接收者。
1、首先避免在每个方法调用中拷贝值（如果值类型是大的结构体的话会更有效率）。
2、其次，方法可以修改接收者指向的值。
尝试修改 Abs 的定义，同时 Scale 方法使用 Vertex 代替 *Vertex 作为接收者。
当 v 是 Vertex 的时候 Scale 方法没有任何作用。`Scale` 修改 `v`。当 v 是一个值（非指针），
方法看到的是 Vertex 的副本，并且无法修改原始值。
Abs 的工作方式是一样的。只不过，仅仅读取 `v`。所以读取的是原始值（通过指针）还是那个值的副本并没有关系。
*/
func test2() {
	v := &Vertex{3, 4}
	v.Scale(5)
	fmt.Println(v, v.Abs())
}
