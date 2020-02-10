package algirithm

import (
	"sort"
)

func canPartition(nums []int) bool {
	sort.Ints(nums)
	sum := 0
	for _, num := range nums {
		sum += num
	}

	if sum&1 == 1 {
		return false
	}

	target := sum / 2

	return dfsCanPartition(nums, len(nums)-1, target, target)
}

func dfsCanPartition(nums []int, idx, had, pass int) bool {
	if had == 0 || pass == 0 {
		return true
	} else if had < 0 || pass < 0 {
		return false
	} else {
		return dfsCanPartition(nums, idx-1, had-nums[idx], pass) || dfsCanPartition(nums, idx-1, had, pass-nums[idx])
	}
}
