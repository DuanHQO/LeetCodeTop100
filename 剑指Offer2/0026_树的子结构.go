package 剑指Offer2

import "fmt"

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}

	var helper func(a, b *TreeNode) bool
	helper = func(a, b *TreeNode) bool {
		if a != nil && b != nil {
			if a.Val != b.Val {
				return false
			}
			return helper(a.Left, b.Left) && helper(a.Right, b.Right)
		}

		return b == nil

	}

	var search func(node *TreeNode) *TreeNode
	search = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}

		if node.Val == B.Val {
			return node
		}

		left := search(node.Left)
		right := search(node.Right)

		if left != nil {
			return left
		}

		if right != nil {
			return right
		}

		return nil
	}

	if search(A) != nil {
		fmt.Println(search(A).Val)
	}
	ret := helper(search(A), B)

	return ret
}
