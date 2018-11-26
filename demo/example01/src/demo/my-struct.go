package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	demo5()
}

func demo5() {
	var a Animal
	s := Student{
		"007",
		"tom",
		23,
	}
	a = &s
	a.Say()
}

func demo4() {
	var p Person
	p.init("0001", "joy", 22)
	s := p.ToString()
	fmt.Println(s)
}

func demo3() {
	var s Student
	s.init("0001", "joy", 11)
	toString := s.ToString()
	fmt.Println(toString)
}

func demo2() {
	var p Person
	p.Id = "0002"
	p.Sex = "man"
	p.int = 123

	fmt.Println(p)
}

func demo1() {

	var stu = Student{
		Id:   "0001",
		Name: "jay",
		Age:  12,
	}

	data, err := json.Marshal(stu)
	if err != nil {
		fmt.Println("json.Marshal() error", err)
	}
	fmt.Println("json.Marshal() success", string(data))

}

// tag
type Student struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// 匿名字段
type Person struct {
	Student
	Sex string
	int
}

// 定义一个动物接口
type Animal interface {
	Say()
}

func (s *Student) Say() {
	fmt.Printf("%v say something...", s.Name)
}

// 方法
func (s *Student) init(id string, name string, age int) {
	s.Id = id
	s.Name = name
	s.Age = age
}

func (s Student) ToString() string {
	return fmt.Sprintf("{id=%v, name=%v, age=%d}", s.Id, s.Name, s.Age)
}
