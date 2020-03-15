package 每日1题3月

import "fmt"

func MaxAreaOfIsland(grid [][]int) int {
	if grid == nil || len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	ret := 0
	cur := 0

	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	fmt.Println(grid)

	var helper func(y, x int)
	helper = func(y, x int) {
		if y < 0 || y == len(grid) || x < 0 || x == len(grid[0]) || grid[y][x] == 0 {
			return
		}

		grid[y][x] = 0
		cur++
		ret = max(ret, cur)
		helper(y-1, x)
		helper(y+1, x)
		helper(y, x-1)
		helper(y, x+1)
	}

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] != 0 {
				cur = 0
				helper(y, x)
			}
		}
	}

	fmt.Println(ret)

	return ret
}
