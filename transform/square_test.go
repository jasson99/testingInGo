package transform

import (
	"reflect"
	"testing"
)

type testSliceItems struct {
	input  []int
	result []int
}

// test case for Square function
func TestTransformSquare(t *testing.T) {
	testDataItem := []testSliceItems{
		{[]int{1, 2, 3, 4, 5}, []int{1, 4, 9, 16, 25}},
		{[]int{0, 0, 0}, []int{0, 0, 0}},
		{[]int{-1, -2, -3, -4, -5, -6}, []int{1, 4, 9, 16, 25, 36}},
	}
	for _, value := range testDataItem {
		result := SquareSlice(value.input)

		if reflect.DeepEqual(result, value.result) {
			t.Logf("SquareSlice PASSED with equal expected and actual results: %v ", result)
		} else {
			t.Errorf("SquareSlice FAILED, expected result was: %v , but got: %v", value.result, result)
		}
	}
}
