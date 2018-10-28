package lesson01

import "fmt"

// channel 是有类型的管道，可以用 channel 操作符 <- 对其发送或者接收值。
func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	// 将和送入c
	// 箭头就是数据流的方向
	c <- sum
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	// 和 map 与 slice 一样，channel 使用前必须创建
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	// 从c中获取
	x, y := <-c, <-c

	fmt.Println(x, y, x+y)
}
