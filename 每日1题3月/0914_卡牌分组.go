package 每日1题3月

import (
	"fmt"
	"sort"
)

func HasGroupsSizeX(deck []int) bool {
	if deck == nil || len(deck) <= 1 {
		fmt.Println(false)
		return false
	}

	sort.Ints(deck)

	cur := 0
	item := deck[0]

	var nums []int
	var gcd func(x, y int) int

	gcd = func(x, y int) int {
		if x == 0 {
			return y
		}
		return gcd(y%x, x)
	}

	for i := 0; i < len(deck); i++ {
		if deck[i] == item {
			cur++
		} else {
			item = deck[i]
			nums = append(nums, cur)
			cur = 1
		}
	}

	nums = append(nums, cur)

	fmt.Println(nums)

	minV := nums[0]

	for _, value := range nums {
		minV = gcd(minV, value)
	}

	fmt.Println(minV >= 2)
	return minV >= 2
}
