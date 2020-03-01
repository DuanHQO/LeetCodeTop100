package 剑指Offer2

func FindRepeatNumber(nums []int) int {
	dic := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		_, ok := dic[nums[i]]
		if !ok {
			dic[nums[i]] = 1
		} else {
			return nums[i]
		}
	}

	return 0
}
