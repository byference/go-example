package utils

import (
	"fmt"
)

func ArrayTest1() {

	var arr [10]int

	arr[0] = 1
	arr[1] = 2

	for index, val := range arr {
		fmt.Printf("[%d]: %d \n", index, val)
	}

}

func ArrayTest2() {

	var arr [2][3]int = [...][3]int{{1, 2, 3}, {4, 5, 6}}

	for row, v1 := range arr {
		for col, v2 := range v1 {
			fmt.Printf("[%d, %d]：%d \n", row, col, v2)
		}
	}

}
