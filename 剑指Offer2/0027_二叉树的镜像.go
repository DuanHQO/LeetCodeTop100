package 剑指Offer2

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	var helper func(node *TreeNode)
	helper = func(node *TreeNode) {
		if node == nil {
			return
		}

		node.Left, node.Right = node.Right, node.Left
		helper(node.Left)
		helper(node.Right)
	}

	helper(root)

	return root
}
