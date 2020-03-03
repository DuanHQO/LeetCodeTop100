package 剑指Offer2

func minArray(numbers []int) int {

	ret := numbers[0]

	for i := 1; i < len(numbers); i++ {
		if numbers[i] < numbers[i-1] {
			ret = numbers[i]
			break
		}
	}

	return ret
}
