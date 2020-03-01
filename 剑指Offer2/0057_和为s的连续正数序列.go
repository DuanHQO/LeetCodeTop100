package 剑指Offer2

import "fmt"

func FindContinuousSequence(target int) [][]int {
	var res [][]int

	for i := 1; i <= target/2; i++ {
		left := target - i
		tmp := []int{i}
		for j := i + 1; i < target && left-j >= 0; j++ {
			tmp = append(tmp, j)
			left -= j
			if left == 0 {
				slice := make([]int, len(tmp), len(tmp))
				copy(slice, tmp)
				res = append(res, slice)
				break
			}
		}
	}

	fmt.Println(res)
	return res
}
