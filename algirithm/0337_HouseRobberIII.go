package algirithm

import (
	"leetcodetop100/kit"
)

func Rob(root *TreeNode) int {
	result := robInternal(root)
	return kit.Max(result[0], result[1])
}

func robInternal(root *TreeNode) []int {
	if root == nil {
		return make([]int, 2)
	}

	result := make([]int, 2)
	left := robInternal(root.Left)
	right := robInternal(root.Right)

	result[0] = kit.Max(left[0], left[1]) + kit.Max(right[0], right[1])
	result[1] = left[0] + right[0] + root.Val

	return result
}
