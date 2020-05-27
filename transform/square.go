package transform

// SquareSlice :return square of a slice
func SquareSlice(s []int) []int {
	squareSlice := make([]int, len(s))

	//for each element in 's', save its square
	for index, value := range s {
		squareSlice[index] = value * value
	}

	//return transformed slice
	return squareSlice
}
