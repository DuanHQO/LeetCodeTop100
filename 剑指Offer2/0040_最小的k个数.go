package 剑指Offer2

import "sort"

func getLeastNumbers(arr []int, k int) []int {
	if arr == nil || k <= 0 {
		return nil
	}

	sort.Ints(arr)
	var res []int

	for i := 0; i < len(arr)-1; i++ {
		res = append(res, arr[i])
		if len(res) == k {
			break
		}
	}

	return res
}
