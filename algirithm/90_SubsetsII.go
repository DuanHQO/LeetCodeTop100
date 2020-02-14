package algirithm

import (
	"sort"
)

func SubsetsWithDup(nums []int) [][]int {
	if nums == nil {
		return nil
	}

	sort.Ints(nums)

	var res [][]int

	var backTrack func(int, []int)

	backTrack = func(i int, tmp []int) {
		if len(tmp) <= len(nums) {
			slice := make([]int, len(tmp))
			copy(slice, tmp)
			res = append(res, slice)
			//return
		}
		for j := i; j < len(nums); j++ {
			if j > i && nums[j] == nums[j-1] {
				continue
			}
			tmp = append(tmp, nums[j])
			backTrack(j+1, tmp)
			tmp = tmp[:len(tmp)-1]
		}
	}

	backTrack(0, []int{})

	return res
}
