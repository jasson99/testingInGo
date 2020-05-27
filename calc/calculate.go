package calc

import (
	"errors"
)

// "CalculateAddition ...""
func CalculateAddition(numbers ...int) (int, error) {
	result := 0

	if len(numbers) < 2 {
		errorMessage := "Please provide at least two numbers for addition"
		return result, errors.New(errorMessage)
	}
	for _, num := range numbers {
		result = result + num
	}

	return result, nil
}
