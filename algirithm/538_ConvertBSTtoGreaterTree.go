package algirithm

import "leetcodetop100/kit"

func convertBST(root *kit.TreeNode) *kit.TreeNode {
	sum := 0
	root = convert(root, &sum)
	return root
}

func convert(root *kit.TreeNode, sum *int) *kit.TreeNode {
	if root == nil {
		return nil
	}
	convert(root.Right, sum)
	root.Val += *sum
	*sum = root.Val
	convert(root.Left, sum)
	return root
}
