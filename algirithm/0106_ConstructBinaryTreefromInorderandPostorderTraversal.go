package algirithm

func buildTreeII(inorder []int, postorder []int) *TreeNode {
	if inorder == nil || postorder == nil {
		return nil
	}

	dic := make(map[int]int)
	for i := 0; i < len(inorder); i++ {
		dic[inorder[i]] = i
	}
	idx := len(postorder) - 1

	var helper func(int, int) *TreeNode
	helper = func(lo int, hi int) *TreeNode {
		if lo == hi {
			return nil
		}

		val := postorder[idx]
		root := &TreeNode{Val: val}
		mid := dic[val]
		idx--

		root.Right = helper(mid+1, hi)
		root.Left = helper(lo, mid-1)

		return root
	}

	root := helper(0, len(postorder)-1)

	return root
}
