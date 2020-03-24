package 每日1题3月

import "fmt"

func Massage(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums)+1)
	dp[0] = 0
	dp[1] = nums[0]

	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := 1; i < len(nums); i++ {
		dp[i+1] = max(nums[i]+dp[i-1], dp[i])
	}

	fmt.Println(dp)

	return dp[len(dp)-1]
}
