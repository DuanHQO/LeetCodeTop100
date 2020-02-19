package algirithm

func plusOne(digits []int) []int {
	var res []int
	if digits == nil {
		return res
	}

	if len(digits) == 0 {
		return []int{1}
	}

	res = make([]int, len(digits), len(digits))
	copy(res, digits)

	res[len(res)-1] += 1

	for i := len(res) - 1; i > 0; i-- {
		if res[i] >= 10 {
			res[i] = res[i] % 10
			res[i-1] += 1
		}
	}

	if res[0] >= 10 {
		res[0] = res[0] % 10
		res = append([]int{1}, res...)
	}

	return res
}
