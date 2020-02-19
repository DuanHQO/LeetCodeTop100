package algirithm

import (
	"fmt"
	"leetcodetop100/kit"
)

func SearchRange(nums []int, target int) []int {
	if nums == nil || len(nums) == 0 {
		return []int{-1, -1}
	}

	length := len(nums)

	var rank func(int) int
	rank = func(key int) int {
		lo := 0
		hi := length - 1

		for lo <= hi {
			mid := lo + (hi-lo)/2
			if nums[mid] < target {
				lo = mid + 1
			} else if nums[mid] == target {
				return mid
			} else {
				hi = mid - 1
			}
		}

		return -1
	}

	idx := rank(target)
	if idx == -1 {
		return []int{-1, -1}
	}

	left := idx
	right := idx

	res := []int{1 << 31, -1 << 31}
	for left >= 0 && nums[left] == target {
		res[0] = kit.Min(res[0], left)
		left--
	}

	for right < length && nums[right] == target {
		res[1] = kit.Max(res[1], right)
		right++
	}

	fmt.Println(res)

	return res
}
