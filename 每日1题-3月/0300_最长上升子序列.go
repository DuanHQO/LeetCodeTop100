package 每日1题_3月

func lengthOfLIS(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return 1
	}

	ret := 0

	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	dp := make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		ret = max(ret, dp[i])
	}

	return ret
}
