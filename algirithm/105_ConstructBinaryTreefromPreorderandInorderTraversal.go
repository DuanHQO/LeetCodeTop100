package algirithm

func buildTree(preorder []int, inorder []int) *TreeNode {
	if preorder == nil || inorder == nil {
		return nil
	}
	if len(preorder) != len(inorder) {
		return nil
	}

	dic := make(map[int]int)
	rootIndex := 0

	for i, v := range inorder {
		dic[v] = i
	}

	var helper func(int, int) *TreeNode
	helper = func(inLeft int, inRight int) *TreeNode {
		if inLeft == inRight {
			return nil
		}

		val := preorder[rootIndex]
		root := &TreeNode{Val: val}
		mid := dic[val]
		rootIndex++

		root.Left = helper(inLeft, mid)
		root.Right = helper(mid+1, inRight)

		return root
	}

	root := helper(0, len(inorder))

	return root
}
