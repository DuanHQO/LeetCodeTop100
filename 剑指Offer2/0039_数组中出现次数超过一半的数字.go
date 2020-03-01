package 剑指Offer2

import (
	"sort"
)

func majorityElement(nums []int) int {
	sort.Ints(nums)
	max := 0
	cur := 0
	idx := 0
	for i := 1; i < len(nums); i++ {
		cur++
		if nums[i] != nums[i-1] {
			cur = 1
		}
		if cur > max {
			idx = i
			max = cur
		}
		if max > len(nums)/2 {
			return nums[i]
		}
	}

	return nums[idx]
}
