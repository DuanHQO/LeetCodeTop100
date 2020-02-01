package main

import (
	"fmt"
	"leetcodetop100/algirithm"
)

func main() {
	grid := [][]int{
		{1, 2, 5},
		{3, 2, 1},
	}
	fmt.Printf("%d", len(grid[0]))
	fmt.Printf("%v\n", algirithm.MinPathSum(grid))
}
