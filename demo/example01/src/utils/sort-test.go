package utils

import (
	"fmt"
	"sort"
)

func TestArraySort() {
	slice := []int{1, 2, 3, 55, 77, 98, 345}

	// 因为数组是值类型，如果排序则修改的是副本，所以这里使用的是引用类型的slice
	sort.Ints(slice)

	fmt.Println("TestArraySort: ", slice)
}

func TestStringSort() {
	slice := []string{"a", "bc", "dea", "zzz", "ngs"}

	// 因为数组是值类型，如果排序则修改的是副本，所以这里使用的是引用类型的slice
	sort.Strings(slice)

	fmt.Println("TestStringSort: ", slice)
}

func TestSearch() {
	slice := []int{1, 2, 3, 55, 77, 98, 345}
	// Search的对象一定是排序后的对象
	index := sort.SearchInts(slice, 77)
	fmt.Println("TestSearch: ", index)

}
