package main

import "fmt"

// Stringer 是一个可以用字符串描述自己的类型。
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("name: %v, age: %v", p.Name, p.Age)
}

func main() {
	p1 := Person{"Joy", 11}
	p2 := Person{"Jay", 11}
	fmt.Println(p1, "; ", p2)
}
