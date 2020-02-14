package algirithm

func subsets(nums []int) [][]int {
	if nums == nil {
		return nil
	}

	var res [][]int

	var backTrack func(int, []int)
	backTrack = func(i int, tmp []int) {
		slice := make([]int, len(tmp), len(tmp))
		copy(slice, tmp)
		res = append(res, slice)
		for j := i; j < len(nums); j++ {
			tmp = append(tmp, nums[j])
			backTrack(j+1, tmp)
			tmp = tmp[:len(tmp)-1]
		}
	}

	backTrack(0, []int{})

	return res
}
