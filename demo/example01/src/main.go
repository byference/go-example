package main

import (
	"fmt"
	"go-example/demo/example01/src/demo"
	"go-example/demo/example01/src/utils"
	"strings"
)

func main() {

	fmt.Println("utils.C", utils.C)
	fmt.Println("utils.Id", utils.Id)
	fmt.Println("utils.Error", utils.Success)
	fmt.Println("utils.Test", utils.Test)
	fmt.Printf("result of utils.Print(): %s.", utils.Print())

	for i := 0; i < 2; i++ {
		j := utils.GetRandomIntn(100)
		fmt.Println("utils.GetRandomInt:", j)
	}

	fmt.Println(utils.Reverse("hello world"))

	demo.StringTest()

	//utils.ArrayTest()
	//utils.MyErrorTest()

	//f := Adder()
	//fmt.Println("Adder-->", f(1))
	//fmt.Println("Adder-->", f(10))
	//fmt.Println("Adder-->", f(100))

	f1 := MakeSuffix(".go")
	fmt.Println(f1("main"))
	fmt.Println(f1("main.go"))

	utils.ArrayTest2()

	utils.SliceTest1()
	utils.SliceTest2()
	utils.SliceTest3()
	utils.SliceTest4()

	utils.TestArraySort()
	utils.TestStringSort()
	utils.TestSearch()

}

// 闭包
func Adder() func(int) int {

	var acc int
	return func(i int) int {
		acc += i
		return acc
	}

}

func MakeSuffix(suffix string) func(string) string {

	return func(name string) string {
		if strings.HasSuffix(name, suffix) == false {
			return name + suffix
		}
		return name
	}

}
