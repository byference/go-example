package lesson02

// var 语句定义了一个变量的列表；跟函数的参数列表一样，类型在后面。
// 就像在这个例子中看到的一样，`var` 语句可以定义在包或函数级别。
import "fmt"

var c, python bool

// 变量定义可以包含初始值，每个变量对应一个
// 如果初始化是使用表达式，则可以省略类型；变量从初始值中获得类型。
var i, j = 1, 2

func main() {
	var o int
	// 在函数中，`:=` 简洁赋值语句在明确类型的地方，可以用于替代 var 定义。
	// 函数外的每个语句都必须以关键字开始（`var`、`func`、等等），`:=` 结构不能使用在函数外。
	a := "this is a string"
	var b string

	fmt.Println(o, c, python)
	fmt.Println(i, j)
	fmt.Println(a, b)
}
