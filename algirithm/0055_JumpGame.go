package algirithm

import "fmt"

func CanJump(nums []int) bool {
	if nums == nil || len(nums) == 0 {
		return false
	}

	lastPos := len(nums) - 1

	for i := len(nums) - 1; i >= 0; i-- {
		if i+nums[i] >= lastPos {
			lastPos = i
		}
	}

	fmt.Println(lastPos == 0)

	return lastPos == 0
	//
	//adj := make(map[int][]int)
	//for i := 0; i < len(nums); i++ {
	//	for j := 0; j <= nums[i]; j++ {
	//		_, ok := adj[i]
	//		if !ok {
	//			adj[i] = []int{}
	//		}
	//		if j + i < len(nums) {
	//			adj[i] = append(adj[i], j + i)
	//		}
	//	}
	//}
	//
	//for i, ints := range adj {
	//	fmt.Printf("%d : %v\n", i, ints)
	//}
	//
	//var dfs func(int)
	//dfs = func(v int)  {
	//	if v >= len(nums) {
	//		return
	//	}
	//	nums[v] *= -1
	//	if v == len(nums) - 1 {
	//		nums[v] = -1
	//		return
	//	} else {
	//		for _, w := range adj[v] {
	//			if w < len(nums) && nums[w] >= 0 {
	//				dfs(w)
	//			}
	//		}
	//	}
	//}
	//
	//dfs(0)
	//fmt.Println(nums)
	//fmt.Println(nums[len(nums) - 1] < 0)
	//
	//return nums[len(nums) - 1] < 0
}
