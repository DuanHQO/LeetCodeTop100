package algirithm

func dailyTemperatures(T []int) []int {
	result := make([]int, len(T))

	for i := 0; i < len(T); i++ {
		for j := i; j < len(T); j++ {
			if T[j] > T[i] {
				result[i] = j - i
				break
			}
		}
	}

	return result
}
