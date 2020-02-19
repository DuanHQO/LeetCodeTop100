package algirithm

import "leetcodetop100/kit"

func MaxProduct(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	dpMax := make([]int, len(nums)+1)
	dpMin := make([]int, len(nums)+1)

	dpMax[0] = 1
	dpMin[0] = 1
	max := -1 << 31
	imax := 1
	imin := 1

	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			imax, imin = imin, imax
		}
		imax = kit.Max(imax*nums[i], nums[i])
		imin = kit.Min(imin*nums[i], nums[i])

		max = kit.Max(imax, max)
	}

	return max
}
