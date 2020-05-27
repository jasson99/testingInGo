package main

import (
	"fmt"
	"testingInGo/calc"
	"testingInGo/transform"
)

func main() {

	s := []int{9, 8, 7}
	square := transform.SquareSlice(s)
	fmt.Println(square)
	fmt.Println("Testing in Go")
	result, error := calc.CalculateAddition(2, 4)
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(result)
	}
}
