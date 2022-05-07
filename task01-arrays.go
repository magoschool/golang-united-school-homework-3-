package homework

func average(input [15]float32) (result float32) {
	var lSum float64 = 0
	var lCount = 0

	for i := 0; i < len(input); i++ {
		lValue := input[i]
		if lValue != 0 {
			lCount++
			lSum += float64(input[i])
		}
	}

	if lCount > 0 {
		return float32(lSum / float64(lCount))
	}

	return 0
}
