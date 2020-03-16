package Tencent

func RemoveDuplicates(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	i := 0
	j := 1
	for ; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	return i + 1
}
