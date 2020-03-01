package 剑指Offer2

import "leetcodetop100/kit"

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var helper func(node *TreeNode) int
	helper = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left := helper(node.Left)
		if left == -1 {
			return -1
		}
		right := helper(node.Right)
		if right == -1 {
			return -1
		}

		if -2 < left-right && left-right < 2 {
			return kit.Max(left, right) + 1
		} else {
			return -1
		}

	}

	return helper(root) != -1
}
