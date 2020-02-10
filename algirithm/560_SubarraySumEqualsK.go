package algirithm

import "fmt"

func SubarraySum(nums []int, k int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	length := len(nums)

	res := 0

	for i := 0; i < length; i++ {
		target := 0
		for j := i; j < length; j++ {
			target += nums[j]
			if target == k {
				res++
				//break
			}
		}
	}

	fmt.Println(res)
	return res
}
