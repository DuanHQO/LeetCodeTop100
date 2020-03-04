package 剑指Offer2

func singleNumbers(nums []int) []int {
	ret := 0
	a, b := 0, 0

	for _, value := range nums {
		ret ^= value
	}

	h := 1
	for (h & ret) == 0 {
		h <<= 1
	}

	for _, value := range nums {
		if h&value == 0 {
			a ^= value
		} else {
			b ^= value
		}
	}

	return []int{a, b}
}
