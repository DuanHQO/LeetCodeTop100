package Tencent

func Search(nums []int, target int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}

	lo := 0
	hi := len(nums) - 1
	mid := lo + (hi-lo)/2

	for lo <= hi {
		if target == nums[mid] {
			return mid
		}
		if nums[lo] <= nums[mid] {
			if nums[lo] <= target && target <= nums[mid] {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		} else {
			if nums[mid] <= target && target <= nums[hi] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
		mid = lo + (hi-lo)/2
	}

	return -1
}
