package utils

import "fmt"

func Reverse(str string) string {
	var result []byte
	tmp := []byte(str)
	for i := len(str); i > 0; i-- {
		result = append(result, tmp[i-1])
	}
	return string(result)
}

func ReverseOld(str string) string {
	var result string
	for i := len(str); i > 0; i-- {
		result = result + fmt.Sprintf("%c", str[i-1])
	}
	return result
}
