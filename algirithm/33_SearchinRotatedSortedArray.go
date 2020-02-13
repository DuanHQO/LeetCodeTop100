package algirithm

func Search(nums []int, target int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}

	left := 0
	right := len(nums) - 1
	mid := left + (right-left)/2

	for left <= right {
		if nums[mid] == target {
			return mid
		}

		if nums[left] <= nums[mid] {
			if nums[left] <= target && target <= nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] <= target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		mid = left + (right-left)/2
	}

	return -1
}
