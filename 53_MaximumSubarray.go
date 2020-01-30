package main

func maxSubArray(nums []int) int {
	if len(nums) <= 1 {
		return nums[0]
	}

	sum, maxSum := -1<<31, -1<<31
	for _, value := range nums {
		sum = max(sum+value, value)
		maxSum = max(maxSum, sum)
	}

	return maxSum
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
