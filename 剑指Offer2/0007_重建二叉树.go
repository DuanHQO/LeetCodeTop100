package 剑指Offer2

func buildTree(preorder []int, inorder []int) *TreeNode {
	if preorder == nil || inorder == nil || len(preorder) != len(inorder) || len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	dic := make(map[int]int)

	for i := 0; i < len(inorder); i++ {
		dic[inorder[i]] = i
	}

	idx := 0

	var helper func(int, int) *TreeNode
	helper = func(lo, hi int) *TreeNode {
		if lo > hi {
			return nil
		}

		val := preorder[idx]
		root := &TreeNode{Val: val}
		mid := dic[val]
		idx++

		left := helper(lo, mid-1)
		right := helper(mid+1, hi)

		root.Left = left
		root.Right = right

		return root
	}

	return helper(0, len(inorder)-1)
}
