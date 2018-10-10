package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

func main() {

	test12()

}

// map 在使用之前必须用 make 而不是 new 来创建；值为 nil 的 map 是空的，并且不能赋值
func test11() {

	// var m map[string]Vertex = make(map[string]Vertex)
	m := make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}

func test12() {

	myMap := make(map[string]string)
	myMap["id"] = "0001"
	myMap["name"] = "old joy"
	myMap["name"] = "joy"
	fmt.Println(myMap)
	// 获取
	fmt.Println(myMap["id"])
	// 删除元素
	delete(myMap, "id")
	fmt.Println(myMap)
	// 通过双赋值检测某个键存在
	stu1, isOk := myMap["name"]
	fmt.Println(stu1, isOk)
	stu2, isOk := myMap["name1"]
	fmt.Println(stu2, isOk)

}
