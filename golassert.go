package golassert

import (
	"fmt"
	"reflect"
	"testing"
)

/*
Type is proxy func for assertType.
*/
func Type(t *testing.T, expected interface{}, result interface{}) {
	assertType(t, true, expected, result)
}

/*
NotType is proxy func for assertType to check no match.
*/
func NotType(t *testing.T, expected interface{}, result interface{}) {
	assertType(t, false, expected, result)
}

/*
Equal is proxy func for assertEqual.
*/
func Equal(t *testing.T, expected interface{}, result interface{}) {
	assertEqual(t, true, expected, result)
}

/*
NotEqual is proxy func for assertEqual to check no match.
*/
func NotEqual(t *testing.T, expected interface{}, result interface{}) {
	assertEqual(t, false, expected, result)
}

/*
assertType asserts type of expected and result.
*/
func assertType(t *testing.T, positive bool, expected interface{}, result interface{}) {
	expectedType := reflect.TypeOf(expected)
	resultType := reflect.TypeOf(result)
	if (expectedType != resultType) == positive {
		err := "Error: Type Assertion"
		t.Errorf("%s\nExpected Value Type: %v\nResult: %v", err, expectedType, resultType)
	}
}

/*
assertEqual asserts if expected result is same as returned result.
*/
func assertEqual(t *testing.T, positive bool, expected interface{}, result interface{}) {
	assertType(t, true, expected, result)
	if expected == nil && result == nil {
		return
	}
	errMsg := fmt.Sprintf("Error: Mismatched Values\n\tExpected value: %v\n\tResult: %v", expected, result)
	switch result.(type) {
	case string, uint, uint8, uint16, uint32, uint64, int, int8, int16, int32, int64, uintptr, float32, float64, complex64, complex128, error, bool:
		if (expected != result) == positive {
			t.Errorf(errMsg)
		}
	default:
		if (fmt.Sprintf("%v", expected) != fmt.Sprintf("%v", result)) == positive {
			t.Errorf(errMsg)
			fmt.Println("[WARN] For non-basic types, might give False -ves for edge cases like diff ordered arrays.")
		}
	}
}
