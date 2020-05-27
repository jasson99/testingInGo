package main

import (
	"fmt"
	"testingInGo/calc"
)

func main() {
	fmt.Println("Testing in Go")
	result, error := calc.CalculateAddition(2, 4)
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(result)
	}
}
