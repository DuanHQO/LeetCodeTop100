package algirithm

import "leetcodetop100/kit"

func maximalSquare(matrix [][]byte) int {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) < 1 {
		return 0
	}

	height := len(matrix)
	width := len(matrix[0])
	maxSide := 0

	dp := make([][]int, height+1)
	for i := 0; i < height; i++ {
		dp[i] = make([]int, width+1)
	}

	for y := 0; y < height+1; y++ {
		for x := 0; x < width; x++ {
			if matrix[y][x] == '1' {
				dp[y+1][x+1] = kit.Min(kit.Min(dp[y+1][x], dp[y][x+1]), dp[y][x]) + 1
				maxSide = kit.Max(maxSide, dp[y+1][x+1])
			}
		}
	}

	return maxSide * maxSide
}
