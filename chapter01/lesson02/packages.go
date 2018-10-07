package lesson02

// 用圆括号组合了导入，这是“打包”导入语句
import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	// 在Go中，首字母大写的名称是被导出的，可以在其他Go文件中使用
	fmt.Println("number is: ", rand.Intn(10))
	fmt.Println("math.Pi is: ", math.Pi)
}
