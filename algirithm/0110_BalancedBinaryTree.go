package algirithm

import "leetcodetop100/kit"

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var helper func(*TreeNode) int
	helper = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		left := helper(root.Left)
		if left == -1 {
			return -1
		}
		right := helper(root.Right)
		if right == -1 {
			return -1
		}

		if -2 < left-right && left-right < 2 {
			return kit.Max(left, right) + 1
		} else {
			return -1
		}

	}

	return helper(root) != 0
}
