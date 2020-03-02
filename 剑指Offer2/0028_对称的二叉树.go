package 剑指Offer2

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var helper func(left, right *TreeNode) bool
	helper = func(left, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}

		if left == nil || right == nil {
			return false
		}

		return left.Val == right.Val && helper(left.Left, right.Right) && helper(left.Right, right.Left)
	}

	return helper(root.Left, root.Right)
}
