package algirithm

import "leetcodetop100/kit"

func minimumTotal(triangle [][]int) int {
	dp := make([]int, len(triangle))

	for i := 0; i < len(triangle); i++ {
		dp[i] = triangle[len(triangle)-1][i]
	}

	for y := len(triangle) - 2; y >= 0; y-- {
		for x := 0; x < len(triangle[y]); x++ {
			dp[x] = triangle[y][x] + kit.Min(dp[x], dp[x+1])
		}
	}

	return dp[0]
}
