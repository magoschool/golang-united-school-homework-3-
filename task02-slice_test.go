package homework

import "testing"

func TestSlice(t *testing.T) {
	var testArray = []int64{1, 2, 5, 15}
	testValue := reverse(testArray)
	var expectedValue = []int64{15, 5, 2, 1}

	if len(testValue) == len(expectedValue) {
		for i := 0; i < len(testValue); i++ {
			if testValue[i] != expectedValue[i] {
				t.Errorf("Invalid array item %v: %v expected but %v found", i, expectedValue[i], testValue[i])
				break
			}
		}
	} else {
		t.Errorf("Invalid array length: %v expected but %v found", len(expectedValue), len(testValue))
	}
}
