package 剑指Offer2

import "fmt"

func TranslateNum(num int) int {

	if num < 10 {
		return 1
	}

	var nums []int
	length := 0
	tmp := num
	for tmp > 0 {
		item := tmp % 10
		nums = append([]int{item}, nums...)
		length++
		tmp /= 10
	}
	nums = append([]int{0}, nums...)

	fmt.Println(nums)
	fmt.Println()
	dp := make([]int, length+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= length; i++ {
		if nums[i-1] > 0 && nums[i-1]*10+nums[i] <= 25 {
			dp[i] = dp[i-1] + dp[i-2]
		} else {
			dp[i] = dp[i-1]
		}
		fmt.Printf("%v\n", dp)
	}

	return dp[length]
}
