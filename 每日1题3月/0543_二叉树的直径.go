package 每日1题3月

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	ret := 0

	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var helper func(node *TreeNode) int
	helper = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left := helper(node.Left)
		right := helper(node.Right)

		ret = max(ret, left+right)
		height := max(left, right) + 1
		return height
	}

	helper(root)

	return ret
}
