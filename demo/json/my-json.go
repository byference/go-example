package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	demo02()
}

// 反序列化
func demo02() {
	var params map[string]interface{}
	params = make(map[string]interface{})
	params["userName"] = "joy"
	params["age"] = 12

	data, err := json.Marshal(params)
	if err != nil {
		fmt.Println("json marshal error")
		return
	}

	str := string(data)
	fmt.Println(str)

	var p map[string]interface{}
	// map这里可以不用make初始化
	err = json.Unmarshal([]byte(str), &p)
	if err != nil {
		fmt.Println("json unmarshal error:", err)
		return
	}
	fmt.Println(p)

}

// 序列化
func demo01() {
	var params map[string]interface{}
	params = make(map[string]interface{})
	params["userName"] = "joy"
	params["age"] = 12

	data, err := json.Marshal(params)
	if err != nil {
		fmt.Println("json marshal error")
		return
	}
	fmt.Println(string(data))

}
