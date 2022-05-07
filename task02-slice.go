package homework

func reverse(input []int64) (result []int64) {
	lLen := len(input)
	lResult := make([]int64, lLen)
	lLen--
	for i := 0; i <= lLen; i++ {
		lResult[i] = input[lLen-i]
	}

	return lResult
}
