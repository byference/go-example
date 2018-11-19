package utils

import "fmt"

func MapTest() {

	// 通过make来创建map
	// var m map[string]string
	// m = make(map[string]string)

	// 建议初始化的时候就给map指定初始容量
	// myMap := make(map[string]string, 10)
	myMap := make(map[string]string)
	// set
	myMap["id"] = "001001"
	myMap["name"] = "joy"
	// get
	id := myMap["id"]
	fmt.Println("MapTest.id", id)
	// delete
	delete(myMap, "id")

	// check,在map嵌套map的时候，赋值前需要判断嵌套的map是否被创建，
	// 如果没有创建则需要make
	s, ok := myMap["id"]
	if ok {
		fmt.Println(s)
	} else {
		fmt.Println("id不存在于myMap")
	}

	// foreach
	for k, v := range myMap {
		fmt.Printf("myMap.foreach-->[%s: %s]\n", k, v)
	}

	fmt.Println("MapTest", myMap)

	myMap1 := map[string]string{
		"username": "admin",
		"password": "admin",
	}
	fmt.Println("myMap1", myMap1)

	myMap2 := make(map[string]map[string]string)
	myMap2["k1"] = make(map[string]string)
	myMap2["k1"]["k1"] = "v1"
	myMap2["k1"]["k2"] = "v2"
	fmt.Println("myMap2", myMap2)

}
