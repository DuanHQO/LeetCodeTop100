package 剑指Offer2

func maxSlidingWindow(nums []int, k int) []int {
	if nums == nil || k <= 0 || k > len(nums) {
		return nil
	}

	var ret []int
	max := -1

	var helper func(int, int) int
	helper = func(l int, r int) int {
		for l < r {
			if nums[l] < nums[r] {
				l++
			} else {
				r--
			}
		}

		return l
	}

	for i := 0; i < len(nums)-k+1; i++ {
		l := i
		r := l + k - 1
		if max == -1 || max == l-1 {
			max = helper(l, r)
		} else {
			if nums[max] < nums[r] {
				max = r
			}
		}
		ret = append(ret, nums[max])
	}

	return ret
}
