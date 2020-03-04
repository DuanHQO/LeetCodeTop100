package 剑指Offer2

func nthUglyNumber(n int) int {
	i2, i3, i5 := 0, 0, 0

	var min = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		dp[i] = min(min(dp[i2]*2, dp[i3]*3), dp[i5]*5)
		if dp[i] == 2*dp[i2] {
			i2++
		}
		if dp[i] == 3*dp[i3] {
			i3++
		}
		if dp[i] == 5*dp[i5] {
			i5++
		}
	}

	return dp[n-1]
}
