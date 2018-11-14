package main

import (
	"fmt"
	"go-example/demo/example01/src/utils"
)

func main() {

	fmt.Println("utils.C", utils.C)
	fmt.Println("utils.Id", utils.Id)
	fmt.Println("utils.Error", utils.Success)
	fmt.Println("utils.Test", utils.Test)
	fmt.Printf("result of utils.Print(): %s.", utils.Print())

	for i := 0; i < 2; i++ {
		j := utils.GetRandomIntn(100)
		fmt.Println("utils.GetRandomInt:", j)
	}

	fmt.Println(utils.Reverse("hello world"))

}
