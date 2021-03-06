package algirithm

import "leetcodetop100/kit"

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums)+1)

	dp[0] = 0
	dp[1] = nums[0]

	for i := 2; i <= len(nums); i++ {
		dp[i] = kit.Max(dp[i-1], dp[i-2]+nums[i-1])
	}

	return dp[len(nums)]
}
