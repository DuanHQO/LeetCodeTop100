package algirithm

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
	}

	for y := 0; y < m; y++ {
		for x := 0; x < n; x++ {
			if obstacleGrid[y][x] == 1 {
				res[y][x] = 0
			} else if x == 0 {
				if y > 0 && res[y-1][x] == 0 {
					res[y][x] = 0
				} else {
					res[y][x] = 1
				}
			} else if y == 0 {
				if x > 0 && res[y][x-1] == 0 {
					res[y][x] = 0
				} else {
					res[y][x] = 1
				}
			} else {
				res[y][x] = res[y-1][x] + res[y][x-1]
			}
		}
	}

	return res[m-1][n-1]
}
