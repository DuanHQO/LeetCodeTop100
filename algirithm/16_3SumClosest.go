package algirithm

import (
	"fmt"
	"sort"
)

func ThreeSumClosest(nums []int, target int) int {

	sort.Ints(nums)
	fmt.Println(nums)
	k := 0
	min := 1 << 31
	offset := 1 << 31

	for ; k < len(nums)-2; k++ {
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}

		i := k + 1
		j := len(nums) - 1

		for i < j {
			sum := nums[k] + nums[i] + nums[j]

			tmp := sum - target

			if tmp < 0 {
				tmp *= -1
			}

			if tmp < offset {
				offset = tmp
				min = sum
			}

			fmt.Printf("nums[k] %d nums[i] %d nums[j] %d\n", nums[k], nums[i], nums[j])
			fmt.Printf("sum %d offset %d\n", sum, tmp)
			fmt.Println()
			if sum == target {
				return sum
			} else if sum > target {
				for i < j && nums[j] == nums[j-1] {
					j--
				}
				j--
			} else {
				for i < j && nums[i] == nums[i+1] {
					i++
				}
				i++
			}

		}
	}

	fmt.Println(min)
	return min
}
