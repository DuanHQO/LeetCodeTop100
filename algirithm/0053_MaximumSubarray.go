package algirithm

import "leetcodetop100/kit"

func maxSubArray(nums []int) int {
	if len(nums) <= 1 {
		return nums[0]
	}

	sum, maxSum := -1<<31, -1<<31
	for _, value := range nums {
		sum = kit.Max(sum+value, value)
		maxSum = kit.Max(maxSum, sum)
	}

	return maxSum
}
