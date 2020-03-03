package 剑指Offer2

import "fmt"

func Search0053(nums []int, target int) int {
	if nums == nil {
		return 0
	}

	ret := 0

	lo := 0
	hi := len(nums) - 1

	for lo <= hi {
		mid := lo + (hi-lo)/2
		fmt.Println(mid)
		if nums[mid] < target {
			lo = mid + 1
		} else if nums[mid] > target {
			hi = mid - 1
		} else {
			ret++
			a := mid - 1
			b := mid + 1

			for a >= 0 {
				if nums[a] == target {
					ret++
					a--
				} else {
					break
				}
			}

			for b <= len(nums)-1 {
				if nums[b] == target {
					ret++
					b++
				} else {
					break
				}
			}

			break
		}
	}

	return ret
}
