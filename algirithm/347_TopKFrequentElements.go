package algirithm

import "sort"

func topKFrequent(nums []int, k int) []int {
	dic := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		_, exists := dic[nums[i]]
		if exists {
			dic[nums[i]] += 1
		} else {
			dic[nums[i]] = 1
		}
	}

	var counts []int
	for _, value := range dic {
		counts = append(counts, value)
	}

	sort.Ints(counts)

	var result []int

	min := counts[len(counts)-k]
	for key, value := range dic {
		if value >= min {
			result = append(result, key)
		}
	}

	return result
}
