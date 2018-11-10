package utils

const (
	Id          = "id0001"
	Success int = 1
)

const (
	A = iota
	B
	C
)

var Test int = 80

func init() {

	Test = 8080 // init方法在全局变量初始化后执行
}

func Print() string {

	return "hello world"
}
