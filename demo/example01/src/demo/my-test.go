package demo

import (
	"fmt"
	"strings"
)

func StringTest() {

	str1 := "hello world"
	str2 := " aa-bb-cc-bb-aa "

	// 判断是否以xx开头
	prefix := strings.HasPrefix(str1, "he")
	fmt.Println("prefix", prefix)

	// 判断是否以xx结尾
	suffix := strings.HasSuffix(str1, "ld")
	fmt.Println("suffix", suffix)

	// 判断字符首次出现的位置，没有则返回-1
	index := strings.Index(str1, "ll")
	fmt.Println("Index", index)

	// 判断字符最后出现的位置，没有则返回-1
	lastIndex := strings.LastIndex(str1, "ll")
	fmt.Println("lastIndex", lastIndex)

	// trim
	space := strings.TrimSpace(str2)
	fmt.Println("space", space)
	strings.TrimLeft(str2, "aa")
	strings.TrimRight(str2, "bb")

	fields := strings.Fields("aa bb cc dd")
	fmt.Println("fields", fields)

	split := strings.Split("aa, bb-cc ,dd", ",")
	fmt.Println("split", split)

}
