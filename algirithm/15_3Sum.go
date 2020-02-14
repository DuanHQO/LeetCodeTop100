package algirithm

import "sort"

func threeSum(nums []int) [][]int {
	if nums == nil || len(nums) == 0 {
		return nil
	}

	sort.Ints(nums)

	if nums[0] > 0 {
		return nil
	}

	k := 0

	var res [][]int

	for ; k < len(nums)-2; k++ {

		if k > 0 && nums[k] == nums[k-1] {
			continue
		}

		i := k + 1
		j := len(nums) - 1

		for i < j {
			sum := nums[k] + nums[i] + nums[j]
			if sum < 0 {
				for i < j && nums[i] == nums[i+1] {
					i++
				}
				i++
			} else if sum > 0 {
				for i < j && nums[j] == nums[j-1] {
					j--
				}
				j--
			} else {
				res = append(res, []int{nums[k], nums[i], nums[j]})
				for i < j && nums[i] == nums[i+1] {
					i++
				}
				i++
				for i < j && nums[j] == nums[j-1] {
					j--
				}
				j--
			}
		}
	}

	return res
}
