package 剑指Offer2

func cuttingRope(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 1

	var max = func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}

	for i := 3; i <= n; i++ {
		for j := 1; j < i; j++ {
			dp[i] = max(max(dp[i], j*(i-j)), j*dp[i-j])
		}
	}

	return dp[n]
}
