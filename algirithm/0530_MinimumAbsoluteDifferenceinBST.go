package algirithm

import "leetcodetop100/kit"

func getMinimumDifference(root *TreeNode) int {
	res := 1 << 31
	var pre *TreeNode

	var inOrder func(node *TreeNode)
	inOrder = func(node *TreeNode) {
		if node == nil {
			return
		}

		inOrder(node.Left)
		if pre != nil {
			res = kit.Min(res, node.Val-pre.Val)
		}
		pre = node
		inOrder(node.Right)
	}

	inOrder(root)

	return res
}
