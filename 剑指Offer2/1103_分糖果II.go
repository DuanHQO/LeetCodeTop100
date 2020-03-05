package 剑指Offer2

import "fmt"

func DistributeCandies(candies int, num_people int) []int {
	ret := make([]int, num_people)
	if candies == 0 {
		return ret
	}

	idx := 0

	for i := 1; candies > 0; i++ {
		idx = idx % num_people
		if candies <= i {
			ret[idx] += candies
			break
		}
		ret[idx] += i
		idx++
		candies -= i
	}

	fmt.Println(ret)

	return ret
}
