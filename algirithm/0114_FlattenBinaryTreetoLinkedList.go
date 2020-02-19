package algirithm

import "leetcodetop100/kit"

type TreeNode = kit.TreeNode

var pre *TreeNode

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	pre = nil
	recur(root)
}

func recur(root *TreeNode) {
	if root == nil {
		return
	}

	recur(root.Right)
	recur(root.Left)
	root.Right = pre
	root.Left = nil
	pre = root
}
