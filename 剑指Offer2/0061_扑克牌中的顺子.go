package 剑指Offer2

import "sort"

func isStraight(nums []int) bool {
	if nums == nil {
		return false
	}

	sort.Ints(nums)
	sum := 0
	for i := 0; i < 4; i++ {
		if nums[i] == 0 {
			continue
		}
		if nums[i] == nums[i+1] {
			return false
		}

		sum += nums[i+1] - nums[i]
	}

	return sum < 5
}
