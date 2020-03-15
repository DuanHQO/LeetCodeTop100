package 每日1题3月

func majorityElement(nums []int) int {
	num := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if num == nums[i] {
			count++
		} else {
			count--
			if count == 0 {
				num = nums[i]
				count = 1
			}
		}
	}

	return num
}
