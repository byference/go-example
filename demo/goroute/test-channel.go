package main

import "fmt"

func main() {

	test2()

}

func test2() {
	pipe := make(chan int, 1)
	go Add(3, 4, pipe)
	sum := <-pipe
	fmt.Println("sum from pipe:", sum)
}

func Add(i1, i2 int, c chan int) int {
	sum := i1 + i2
	c <- sum
	return sum
}

func test1() {
	pipe := make(chan int, 3)
	pipe <- 1
	pipe <- 2
	pipe <- 3
	var t1 int
	t1 = <-pipe
	fmt.Println(t1)
}
