package algirithm

import (
	"leetcodetop100/kit"
)

func NthUglyNumber(n int) int {
	dp := make([]int, n)
	dp[0] = 1
	//三指针算法
	dp2, dp3, dp5 := 0, 0, 0
	for i := 1; i < n; i++ {
		dp[i] = kit.Min(kit.Min(2*dp[dp2], 3*dp[dp3]), 5*dp[dp5])
		if dp[i] == 2*dp[dp2] {
			dp2++
		}
		if dp[i] == 3*dp[dp3] {
			dp3++
		}
		if dp[i] == 5*dp[dp5] {
			dp5++
		}
	}
	return dp[len(dp)-1]
}
