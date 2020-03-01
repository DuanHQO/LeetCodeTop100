package 剑指Offer2

func exchange(nums []int) []int {
	if nums == nil {
		return nil
	}

	var left []int
	var right []int

	for i := 0; i < len(nums); i++ {
		if nums[i]&1 == 1 {
			left = append(left, nums[i])
		} else {
			right = append(right, nums[i])
		}
	}

	return append(left, right...)
}
