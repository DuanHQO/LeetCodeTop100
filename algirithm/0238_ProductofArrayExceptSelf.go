package algirithm

func productExceptSelfA(nums []int) []int {
	if nums == nil || len(nums) <= 1 {
		return nil
	}

	left := make([]int, len(nums))
	right := make([]int, len(nums))

	left[0] = 1
	for i := 1; i < len(nums); i++ {
		left[i] = left[i-1] * nums[i-1]
	}

	right[len(nums)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i+1]
	}

	result := make([]int, len(nums))
	for i, _ := range nums {
		result[i] = left[i] * right[i]
	}

	return result
}

func productExceptSelfB(nums []int) []int {
	if nums == nil || len(nums) <= 1 {
		return nil
	}

	result := make([]int, len(nums))
	result[0] = 1
	for i := 1; i < len(nums); i++ {
		result[i] = result[i-1] * nums[i-1]
	}

	r := 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] = result[i] * r
		r *= nums[i]
	}

	return result
}
