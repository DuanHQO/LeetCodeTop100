package algirithm

import "leetcodetop100/kit"

func maxProfitIV(k int, prices []int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	}

	maxK := k
	dp := make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, maxK+1)
		dp[i][0] = make([]int, 2)
		for k := maxK; k >= 1; k-- {
			dp[i][k] = make([]int, 2)
			if i-1 == -1 {
				dp[i][k][0] = 0
				dp[i][k][1] = -prices[i]
				continue
			}
			dp[i][k][0] = kit.Max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
			dp[i][k][1] = kit.Max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])
		}
	}
	return dp[n-1][maxK][0]
}
