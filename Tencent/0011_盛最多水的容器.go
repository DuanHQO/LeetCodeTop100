package Tencent

import "fmt"

func MaxArea(height []int) int {
	lo := 0
	hi := len(height) - 1

	ret := 0

	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var min = func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	for lo < hi {
		small := min(height[lo], height[hi])
		ret = max(ret, (hi-lo)*small)
		if height[lo] < height[hi] {
			lo++
		} else {
			hi--
		}
	}

	fmt.Println(ret)

	return ret
}
