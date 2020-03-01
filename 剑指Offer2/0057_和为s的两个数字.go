package 剑指Offer2

func twoSum(nums []int, target int) []int {
	if nums == nil || nums[0] > target {
		return nil
	}

	hi := len(nums) - 1
	lo := 0

	for lo < hi {
		if nums[lo]+nums[hi] > target {
			hi--
		} else if nums[lo]+nums[hi] < target {
			lo++
		} else {
			return []int{nums[lo], nums[hi]}
		}
	}

	return nil
}
