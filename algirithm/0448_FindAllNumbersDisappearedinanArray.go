package algirithm

func findDisappearedNumbers(nums []int) []int {
	var slice []int
	for i := 0; i < len(nums); i++ {
		index := nums[i]
		if index < 0 {
			index *= -1
		}
		if nums[index-1] > 0 {
			nums[index-1] *= -1
		}
	}

	for index, value := range nums {
		if value > 0 {
			slice = append(slice, index+1)
		}
	}
	return slice
}
