package algirithm

import "leetcodetop100/kit"

func lengthOfLIS(nums []int) int {
	if nums == nil || len(nums) < 1 {
		return 0
	}

	if len(nums) <= 1 {
		return 1
	}

	dp := make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}

	res := 0

	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = kit.Max(dp[i], dp[j]+1)
			}
		}
		res = kit.Max(res, dp[i])
	}

	return res
}
