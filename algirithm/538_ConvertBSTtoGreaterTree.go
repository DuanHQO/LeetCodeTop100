package algirithm

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	root = convert(root, &sum)
	return root
}

func convert(root *TreeNode, sum *int) *TreeNode {
	if root == nil {
		return nil
	}
	convert(root.Right, sum)
	root.Val += *sum
	*sum = root.Val
	convert(root.Left, sum)
	return root
}
