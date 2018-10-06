package main

import (
	"go-example/chapter01/lesson02/util"
	"time"
)

func init() {
	util.Println("init...,this time is:", time.Now())
}

func main() {
	msg := "hello,world"
	util.Println(msg)

	sum := util.Add(3, 4)
	util.Println("求和结果：", sum)

}
