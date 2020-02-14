package algirithm

func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	return helperPathSum(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}

func helperPathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}

	res := 0
	if root.Val == sum {
		res++
	}
	res += helperPathSum(root.Left, sum-root.Val)
	res += helperPathSum(root.Right, sum-root.Val)
	return res
}
