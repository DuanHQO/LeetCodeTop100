package algirithm

import (
	"leetcodetop100/kit"
)

func Rob(root *TreeNode) int {

	var helper func(node *TreeNode) []int
	helper = func(node *TreeNode) []int {
		if node == nil {
			return make([]int, 2)
		}

		result := make([]int, 2)
		left := helper(node.Left)
		right := helper(node.Right)

		result[0] = kit.Max(left[0], left[1]) + kit.Max(right[0], right[1])
		result[1] = left[0] + right[0] + node.Val

		return result
	}

	res := helper(root)

	return kit.Max(res[0], res[1])
}
