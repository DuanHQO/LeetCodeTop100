package algirithm

func NumIslands(grid [][]byte) int {
	res := 0

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == '1' {
				dfsNumIslands(y, x, grid)
				res++
			}
		}
	}

	return res
}

func dfsNumIslands(y, x int, grid [][]byte) {
	ny := len(grid)
	nx := len(grid[0])

	if y < 0 || x < 0 || y >= ny || x >= nx || grid[y][x] == '0' {
		return
	}

	grid[y][x] = '0'
	dfsNumIslands(y-1, x, grid)
	dfsNumIslands(y+1, x, grid)
	dfsNumIslands(y, x-1, grid)
	dfsNumIslands(y, x+1, grid)
}
