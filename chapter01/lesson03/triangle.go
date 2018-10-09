package main

import "fmt"

func main() {

	PrintTriangle(6)
}

func PrintTriangle(size int) {

	for x := 1; x < size+1; x++ {
		for y := size; y > x; y-- {
			fmt.Print(" ")
		}
		for z := x; z > 0; z-- {
			fmt.Print("*")
		}
		fmt.Println("")
	}

}
