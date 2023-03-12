package golassert

import (
	"fmt"
	"reflect"
	"testing"
)

/*
Type is proxy func for AssertType.
*/
func Type(t *testing.T, expected interface{}, result interface{}) {
	AssertType(t, expected, result)
}

/*
Equal is proxy func for AssertEqual.
*/
func Equal(t *testing.T, expected interface{}, result interface{}) {
	AssertEqual(t, expected, result)
}

/*
AssertType asserts type of expected and result.
*/
func AssertType(t *testing.T, expected interface{}, result interface{}) {
	expectedType := reflect.TypeOf(expected)
	resultType := reflect.TypeOf(result)
	if expectedType != resultType {
		err := "Error: [AssertEqual] Mismatched Types"
		t.Errorf("%s\nExpected Value Type: %v\nResult: %v", err, expectedType, resultType)
	}
}

/*
AssertEqual asserts if expected result is same as returned result.
*/
func AssertEqual(t *testing.T, expected interface{}, result interface{}) {
	AssertType(t, expected, result)
	if expected == nil && result == nil {
		return
	}
	errMsg := fmt.Sprintf("Error: [] Mismatched Values\nExpected value: %v\nResult: %v", expected, result)
	switch result.(type) {
	case string, uint, uint8, uint16, uint32, uint64, int, int8, int16, int32, int64, uintptr, float32, float64, complex64, complex128, error, bool:
		if expected != result {
			t.Errorf(errMsg)
		}
	default:
		if fmt.Sprintf("%v", expected) != fmt.Sprintf("%v", result) {
			t.Errorf(errMsg)
		}
	}

}
