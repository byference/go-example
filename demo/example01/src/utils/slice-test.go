package utils

import "fmt"

func SliceTest() {
	var arr []int
	arr = append(arr, 1, 2)
	arr = append(arr, arr...)
	fmt.Println(arr)
}

// 通过已存在的数组创建slice
func SliceTest1() {

	arr := [...]int{1, 2, 3, 4}
	var s []int
	s = arr[:]
	fmt.Println("SliceTest1: ", s)

}

// 通过make创建
func SliceTest2() {

	arr := []int{1, 2, 3, 4}

	var s []int = make([]int, 5)
	// s = append(s, 1,2,3)
	copy(s, arr)
	fmt.Println("SliceTest2: ", s)

}

func SliceTest3() {

	str := "hello world"
	slice := str[:]
	for _, val := range slice {
		fmt.Printf("SliceTest3: %c\n", val)
	}

}

func SliceTest4() {

	str := "你好， 世界"
	r := []rune(str)
	r[0] = '0'
	s := string(r)
	fmt.Println(s)

}
