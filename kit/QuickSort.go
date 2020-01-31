package kit

func QuickSort(nums []int, start, end int) {
	if end <= start {
		return
	}

	j := partition(nums, start, end)
	QuickSort(nums, start, j-1)
	QuickSort(nums, j+1, end)
}

func partition(nums []int, start int, end int) (result int) {
	i := start
	j := end + 1
	cValue := nums[start]

	for j := i + 1; j < end; j++ {
		if nums[j] < cValue {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i], nums[start] = nums[start], nums[i]

	return i
}

func Quick3Way(nums []int, lo, hi int) {
	if hi <= lo {
		return
	}

	slow := lo
	fast := lo + 1
	gt := hi
	cValue := nums[lo]

	for fast <= gt {
		if nums[fast] < cValue {
			nums[fast], nums[slow] = nums[slow], nums[fast]
			fast++
			slow++
		} else if nums[fast] > cValue {
			nums[fast], nums[gt] = nums[gt], nums[fast]
			gt--
		} else {
			fast++
		}
	}
	Quick3Way(nums, lo, slow-1)
	Quick3Way(nums, fast, hi)
}
