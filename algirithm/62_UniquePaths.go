package algirithm

func uniquePaths(m int, n int) int {
	grid := make([][]int, m)

	for i := 0; i < len(grid); i++ {
		grid[i] = make([]int, n)
	}

	for y := 0; y < m; y++ {
		for x := 0; x < n; x++ {
			if y == 0 && x == 0 {
				grid[y][x] = 1
			} else if y == 0 {
				grid[y][x] = 1
			} else if x == 0 {
				grid[y][x] = 1
			} else {
				grid[y][x] = grid[y-1][x] + grid[y][x-1]
			}
		}
	}

	return grid[m-1][n-1]
}
