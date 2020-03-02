package 剑指Offer2

import "leetcodetop100/kit"

func maxSubArray(nums []int) int {
	if len(nums) <= 1 {
		return nums[0]
	}

	sum, max := -1<<31, -1<<31
	for _, value := range nums {
		sum = kit.Max(sum+value, value)
		max = kit.Max(max, sum)
	}

	return max
}
