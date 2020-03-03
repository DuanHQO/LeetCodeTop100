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
		//如果滑出去的是最大值，就在当前窗口内重新更新
		if max == -1 || max == l-1 {
			max = helper(l, r)
		} else {
			//如果滑出去的不是最大值，就跟新进来的比较
			if nums[max] < nums[r] {
				max = r
			}
		}
		ret = append(ret, nums[max])
	}

	return ret
}
