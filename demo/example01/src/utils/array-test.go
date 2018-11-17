package utils

import "fmt"

func ArrayTest() {
	var arr []int
	arr = append(arr, 1, 2)
	arr = append(arr, arr...)
	fmt.Println(arr)
}
