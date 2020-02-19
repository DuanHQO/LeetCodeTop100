package algirithm

import "fmt"

func FindTargetSumWays(nums []int, S int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	res := 0
	dfsFindTargetSumWays(nums, S, 0, 0, &res)

	fmt.Println(res)

	return res
}

func dfsFindTargetSumWays(nums []int, target, idx int, sum int, count *int) {
	if idx == len(nums) {
		if sum == target {
			*count++
		}
		return
	}

	dfsFindTargetSumWays(nums, target, idx+1, sum+nums[idx], count)
	dfsFindTargetSumWays(nums, target, idx+1, sum-nums[idx], count)
}
