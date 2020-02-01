package algirithm

import (
	"leetcodetop100/kit"
)

func MinPathSum(grid [][]int) int {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if x == 0 && y == 0 {
				grid[y][x] = grid[y][x]
			} else if x != 0 && y != 0 {
				grid[y][x] = kit.Min(grid[y-1][x], grid[y][x-1]) + grid[y][x]
			} else if x == 0 {
				grid[y][x] = grid[y-1][x] + grid[y][x]
			} else if y == 0 {
				grid[y][x] = grid[y][x-1] + grid[y][x]
			}
		}
	}

	return grid[len(grid)-1][len(grid[0])-1]
}
