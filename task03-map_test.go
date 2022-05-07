package homework

import "testing"

func testMapOutput(t *testing.T, aInput map[int]string, aResult []string) {
	testValue := sortMapValues(aInput)

	if len(testValue) == len(aResult) {
		for i := 0; i < len(testValue); i++ {
			if testValue[i] != aResult[i] {
				t.Errorf("Invalid map item %v: '%v' expected but '%v' found", i, aResult[i], testValue[i])
				break
			}
		}
	} else {
		t.Errorf("Invalid map length: %v expected but %v found", len(aResult), len(testValue))
	}
}

func TestMap1(t *testing.T) {
	/*
	   Input -> {2: "a", 0: "b", 1: "c"}
	   Output -> ["b", "c", "a"]
	*/
	lTestValue := map[int]string{2: "a", 0: "b", 1: "c"}
	lResult := []string{"b", "c", "a"}
	testMapOutput(t, lTestValue, lResult)
}

func TestMap2(t *testing.T) {
	/*
	   Input -> {10: "aa", 0: "bb", 500: "cc"}
	   Output -> ["bb", "aa", "cc"]
	*/
	lTestValue := map[int]string{10: "aa", 0: "bb", 500: "cc"}
	lResult := []string{"bb", "aa", "cc"}
	testMapOutput(t, lTestValue, lResult)
}
