package algirithm

var dic map[int]int
var thisPre []int
var rootIndex int

func buildTree(preorder []int, inorder []int) *TreeNode {
	if preorder == nil || inorder == nil {
		return nil
	}
	if len(preorder) != len(inorder) {
		return nil
	}

	rootIndex = 0
	dic = make(map[int]int)
	for i, v := range inorder {
		dic[v] = i
	}

	thisPre = preorder
	root := helperBuildTree(0, len(inorder))

	return root
}

func helperBuildTree(inLeft int, inRight int) *TreeNode {
	if inLeft == inRight {
		return nil
	}

	val := thisPre[rootIndex]
	root := &TreeNode{Val: val}
	mid := dic[val]
	rootIndex++

	root.Left = helperBuildTree(inLeft, mid)
	root.Right = helperBuildTree(mid+1, inRight)
	return root
}
