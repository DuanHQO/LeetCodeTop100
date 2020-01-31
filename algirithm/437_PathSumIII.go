package algirithm

func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	return hasSum(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}

func sumGo(root *TreeNode, sum int, count *int) {
	if root == nil {
		return
	}
	sum -= root.Val
	if sum == 0 {
		*count++
	}
	sumGo(root.Left, sum, count)
	sumGo(root.Right, sum, count)
}

func hasSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}

	sum -= root.Val

	if sum == 0 {
		return 1 + hasSum(root.Left, sum) + hasSum(root.Right, sum)
	} else {
		return hasSum(root.Left, sum) + hasSum(root.Right, sum)
	}
}
