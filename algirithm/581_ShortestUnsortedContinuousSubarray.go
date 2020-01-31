package algirithm

func findUnsortedSubarray(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}
	left, right := 0, -1
	min, max := nums[n-1], nums[0]

	for i := 1; i < n; i++ {
		if max <= nums[i] {
			max = nums[i]
		} else {
			right = i
		}

		j := n - i - 1
		if nums[j] <= min {
			min = nums[j]
		} else {
			left = j
		}
	}

	return right - left + 1
}
