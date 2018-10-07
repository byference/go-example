package lesson02

// 函数可以没有参数或接受多个参数
import "fmt"

func main() {

	c := add(1, 2)
	fmt.Println(c)

	fmt.Println(swap("hello", "world"))

	fmt.Println(test(123))

}

// 当两个或多个连续的函数命名参数是同一类型，则除了最后一个类型之外，其他都可以省略
func add(x, y int) int {
	return x + y
}

// 函数可以返回任意数量的返回值
func swap(x, y string) (string, string) {
	return y, x
}

/*
	Go 的返回值可以被命名，并且像变量那样使用。
	返回值的名称应当具有一定的意义，可以作为文档使用。
	没有参数的 return 语句返回结果的当前值。也就是`直接`返回。
	直接返回语句仅应当用在像下面这样的短函数中。在长的函数中它们会影响代码的可读性。
*/
func test(sum int) (x, y int) {
	x = 1
	y = 2
	return
}
