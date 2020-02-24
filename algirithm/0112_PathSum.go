package algirithm

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	var helper func(int, *TreeNode) bool
	helper = func(tmpSum int, root *TreeNode) bool {
		if root == nil {
			return false
		}

		tmpSum += root.Val

		if tmpSum == sum && root.Left == nil && root.Right == nil {
			return true
		}

		return helper(tmpSum, root.Left) || helper(tmpSum, root.Right)

	}

	return helper(0, root)
}
