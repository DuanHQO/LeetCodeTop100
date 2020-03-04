package 剑指Offer2

func maxValue(grid [][]int) int {
	if grid == nil || len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	dp := make([][]int, len(grid))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(grid[0]))
	}

	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if x == 0 && y == 0 {
				dp[y][x] = grid[y][x]
			} else if y == 0 {
				dp[y][x] = dp[y][x-1] + grid[y][x]
			} else if x == 0 {
				dp[y][x] = dp[y-1][x] + grid[y][x]
			} else {
				dp[y][x] = max(dp[y-1][x], dp[y][x-1]) + grid[y][x]
			}
		}
	}

	return dp[len(grid)-1][len(grid[0])-1]
}
