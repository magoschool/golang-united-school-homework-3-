package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	// sort keys keys
	lKeys := make([]int, 0, len(input))
	for lKey := range input {
		lKeys = append(lKeys, lKey)
	}
	sort.Ints(lKeys)

	// init result
	lResult := make([]string, 0, len(input))
	for _, lKey := range lKeys {
		lValue := input[lKey]
		lResult = append(lResult, lValue)
	}

	return lResult
}
