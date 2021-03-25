package twosums

import (
	"reflect"
	"testing"
)

func TestTwoSumsFirstTwo(t *testing.T) {
	givenNumbers := []int{1, 2, 3, 5, 8}
	var givenSum int = 7

	returnedValue := twoSum(givenNumbers, givenSum)
	expectedValue := []int{1, 3}

	if !reflect.DeepEqual(expectedValue, returnedValue) {
		t.Errorf("Wanted [1 3], got %v", returnedValue)
	}

}
