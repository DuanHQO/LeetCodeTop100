package algirithm

func deleteNode(root *TreeNode, key int) *TreeNode {
	var deleteMin func(*TreeNode) *TreeNode
	deleteMin = func(node *TreeNode) *TreeNode {
		if node.Left == nil {
			return node.Right
		}
		node.Left = deleteMin(node.Left)
		return node
	}

	var min func(*TreeNode) *TreeNode
	min = func(node *TreeNode) *TreeNode {
		if node.Left == nil {
			return node
		}
		return min(node.Left)
	}

	var deleteKey func(*TreeNode, int) *TreeNode
	deleteKey = func(node *TreeNode, key int) *TreeNode {
		if node == nil {
			return nil
		}

		if key > node.Val {
			node.Right = deleteKey(node.Right, key)
		} else if key < node.Val {
			node.Left = deleteKey(node.Left, key)
		} else {
			if node.Right == nil {
				return node.Left
			}
			if node.Left == nil {
				return node.Right
			}

			t := node
			node = min(t.Right)
			node.Right = deleteMin(t.Right)
			node.Left = t.Left
		}
		return node
	}

	root = deleteKey(root, key)

	return root
}
