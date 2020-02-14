package algirithm

import "sort"

func permuteUnique(nums []int) [][]int {
	if nums == nil {
		return nil
	}

	sort.Ints(nums)
	var res [][]int
	marked := make([]bool, len(nums))

	var backTrack func([]int)

	backTrack = func(tmp []int) {
		if len(tmp) == len(nums) {
			slice := make([]int, len(nums), len(nums))
			copy(slice, tmp)
			res = append(res, slice)
		} else {
			for i := 0; i < len(nums); i++ {
				if marked[i] {
					continue
				}
				if i > 0 && nums[i] == nums[i-1] && marked[i-1] {
					continue
				}
				marked[i] = true
				tmp = append(tmp, nums[i])
				backTrack(tmp)
				marked[i] = false
				tmp = tmp[:len(tmp)-1]
			}
		}
	}

	backTrack([]int{})

	return res
}
