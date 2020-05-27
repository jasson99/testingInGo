package transform

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
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

func TestReadFile(t *testing.T) {
	data, err := ioutil.ReadFile("../testData/test.data")
	if err != nil {
		t.Errorf("The test.data file could not be opened.")
	}
	if string(data) != "hello world" {
		t.Errorf("String contents donot match as expected. The actual string is %s", string(data))
	}
}

func TestHttpRequest(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "{ \"status\": \"good\"}")
	}

	req := httptest.NewRequest("GET", "https://tutorialedge.net", nil)
	w := httptest.NewRecorder()
	handler(w, req)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	if 200 != resp.StatusCode {
		t.Errorf("Status Code Not OK")
	}
}
