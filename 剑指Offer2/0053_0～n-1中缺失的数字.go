package 剑指Offer2

func missingNumber(nums []int) int {
	ret := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] > i {
			ret = i
			break
		}
	}

	return ret
}
