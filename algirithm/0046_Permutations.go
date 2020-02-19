package algirithm

func permute(nums []int) [][]int {
	if nums == nil {
		return nil
	}

	marked := make([]bool, len(nums))
	var res [][]int

	var backTrack func([]int)

	backTrack = func(tmp []int) {
		if len(tmp) == len(nums) {
			slice := make([]int, len(tmp), len(tmp))
			copy(slice, tmp)
			res = append(res, slice)
			return
		}
		for j := 0; j < len(nums); j++ {
			if marked[j] {
				continue
			}
			marked[j] = true
			tmp = append(tmp, nums[j])
			backTrack(tmp)
			marked[j] = false
			tmp = tmp[:len(tmp)-1]
		}
	}

	backTrack([]int{})

	return res
}
