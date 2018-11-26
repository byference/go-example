package main

import (
	"fmt"
	"math"
	"math/rand"
)

var code = []string{
	"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j",
	"k", "l", "m", "n", "o", "p", "q", "r", "s", "t",
	"u", "v", "w", "x", "y", "z",
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J",
	"K", "L", "M", "N", "O", "P", "Q", "R", "S", "T",
	"U", "V", "W", "X", "Y", "Z",
	"-", "_",
}

// 十进制转64进制
// 转换其他进制同理
func Trans(value int) string {

	number := 64
	if number > len(code) {
		// 返回一个异常
		return "error"
	}
	// 负数处理
	if value < 0 {
		return "-" + Trans(0-value)
	}
	if value < number {
		return code[value]
	} else {
		n := value % number
		return Trans(value/number) + code[n]
	}
}

// 获取一个较大的随机数
// 实际上有一个提供自增id的发号器来提供
func getRandInt() int {
	return rand.Intn(math.MaxInt32)
}

func main() {
	n := Trans(getRandInt())
	fmt.Println("--> main：", n)
}
