package main

import (
	"fmt"
	"strconv"
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

	// 将数组用符号连接成字符串
	strs := []string{"a", "b", "c"}
	join := strings.Join(strs, ",")
	fmt.Println("join", join)

	// 重复3次qwe
	repeat := strings.Repeat("qwe", 3)
	fmt.Println("repeat", repeat)

	// 字符串替换，替换2次
	replace := strings.Replace("aabbcc", "bb", "--", 2)
	fmt.Println("replace", replace)

	// 统计字符出现次数
	count := strings.Count("aabbcaaca", "aa")
	fmt.Println("count", count)

	// 数字转字符
	itoa := strconv.Itoa(1000)
	fmt.Println("itoa", itoa)

	// 字符转数字
	num, err := strconv.Atoi("100")
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println("num", num)

}
