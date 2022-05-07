package homework

import "testing"

func TestAvg(t *testing.T) {
	var testArray = [15]float32{1, 2, 3, 4, 5, 6}
	testValue := average(testArray)
	var expectedValue float32 = 3.5

	if testValue != expectedValue {
		t.Errorf("Array average failed: %v expected but %v found", expectedValue, testValue)
	}
}
