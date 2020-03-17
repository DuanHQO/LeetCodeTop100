package Tencent

import "fmt"

func Permute(nums []int) [][]int {
	var ret [][]int
	if nums == nil || len(nums) == 0 {
		return ret
	}

	marked := make([]bool, len(nums))
	length := len(nums)

	var helper func([]int)
	helper = func(ints []int) {
		if len(ints) == length {
			slice := make([]int, length)
			copy(slice, ints)
			ret = append(ret, slice)
			return
		}

		for i := 0; i < length; i++ {
			if marked[i] {
				continue
			}
			marked[i] = true
			ints = append(ints, nums[i])
			helper(ints)
			marked[i] = false
			ints = ints[:len(ints)-1]
		}
	}

	helper([]int{})

	for _, value := range ret {
		fmt.Println(value)
	}

	return ret
}
