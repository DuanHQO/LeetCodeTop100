package algirithm

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	if candidates == nil {
		return nil
	}

	var res [][]int

	sort.Ints(candidates)

	if len(candidates) == 0 || candidates[0] > target {
		return res
	}

	var dfs func(int, int, []int)
	dfs = func(idx, left int, tmp []int) {
		if left == 0 {
			slice := make([]int, len(tmp), len(tmp))
			copy(slice, tmp)
			res = append(res, slice)
			return
		}

		for i := idx; i < len(candidates) && left-candidates[i] >= 0; i++ {
			if i > idx && candidates[i] == candidates[i-1] {
				continue
			}
			tmp = append(tmp, candidates[i])
			dfs(i+1, left-candidates[i], tmp)
			tmp = tmp[:len(tmp)-1]
		}

	}

	dfs(0, target, []int{})

	return res
}
