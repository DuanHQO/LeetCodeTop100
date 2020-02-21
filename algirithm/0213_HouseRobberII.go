package algirithm

import "leetcodetop100/kit"

func robII(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	var rob func([]int) int
	rob = func(nums []int) int {
		pre := 0
		cur := 0
		tmp := 0
		for i := 0; i < len(nums); i++ {
			tmp = cur
			cur = kit.Max(pre+nums[i], cur)
			pre = tmp
		}
		return cur
	}

	return kit.Max(rob(nums[:len(nums)-1]), rob(nums[1:]))
}
