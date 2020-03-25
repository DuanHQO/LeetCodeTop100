package 每日1题3月

import "fmt"

func SurfaceArea(grid [][]int) int {
	if grid == nil || len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	area := 0

	var min = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] > 0 {
				area += grid[y][x]<<2 + 2
				if y > 0 {
					area -= min(grid[y][x], grid[y-1][x]) * 2
				}
				if x > 0 {
					area -= min(grid[y][x], grid[y][x-1]) * 2
				}
			}
		}
	}

	fmt.Println(area)

	return area
}
