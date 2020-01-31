package algirithm

import (
	"sort"
)

func CombinationSum(candidates []int, target int) [][]int {
	if candidates == nil {
		return nil
	}

	var result [][]int

	if len(candidates) == 0 || target < 0 {
		return result
	}

	sort.Ints(candidates)

	dfs39(candidates, target, 0, &result, []int{})

	return result
}

func dfs39(candidates []int, left int, start int, result *[][]int, pre []int) {
	temp := make([]int, len(pre))
	for i, v := range pre {
		temp[i] = v
	}
	if left < 0 {
		return
	}

	if left == 0 {
		*result = append(*result, temp)
		return
	}

	for i := start; i < len(candidates) && left-candidates[i] >= 0; i++ {
		temp = append(temp, candidates[i])
		dfs39(candidates, left-candidates[i], i, result, temp)
		temp = temp[:len(temp)-1]
	}

}
