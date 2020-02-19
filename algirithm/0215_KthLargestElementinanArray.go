package algirithm

import (
	"sort"
)

func FindKthLargest(nums []int, k int) int {
	if nums == nil {
		return -1
	}

	sort.Ints(nums)
	cur := nums[len(nums)-k]

	return cur
}
