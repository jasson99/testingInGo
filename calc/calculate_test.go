package calc

import (
	"testing"
)

type TestDataItem struct {
	input    []int
	result   int
	hasError bool
}

func TestCalculateAddition(t *testing.T) {
	dataItems := []TestDataItem{
		{[]int{1, 2, 3, 4}, 10, false},
		{[]int{2, 4, -1}, 5, false},
		{[]int{9999, 0}, 9999, false},
		{[]int{9}, 0, true},
	}

	for _, dataItem := range dataItems {
		result, err := CalculateAddition(dataItem.input...)

		if dataItem.hasError {
			if err == nil {
				t.Errorf("calculateAddition with arguments %d : FAILED . An error was expected but got result: %d", dataItem.input, result)
			} else if result != dataItem.result {
				t.Errorf("calculateAddition with arguments %d : FAILED . The expected result is : %d , but got : %d instead.", dataItem.input, dataItem.result, result)
			} else {
				t.Logf("calculateAddition with arguments %v : PASSED. Expected an error and got an error : %v", dataItem.input, err)
			}
		} else {
			if err != nil {
				t.Errorf("calculateAddition with arguments %d : FAILED . An unexpected error is found : %d", dataItem.input, err)
			} else {
				if dataItem.result != result {
					t.Errorf("calculateAddition with arguments %d : FAILED . The expected result is : %d , but got : %d instead.", dataItem.input, dataItem.result, result)
				} else {
					t.Logf("calculateAddition with arguments %v : PASSED. The expected result was : %v, and got : %v", dataItem.input, dataItem.result, result)

				}
			}
		}
	}
}
