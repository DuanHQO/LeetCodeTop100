package algirithm

import "strconv"

func binaryTreePaths(root *TreeNode) []string {
	var res []string
	if root == nil {
		return res
	}

	var dfs func(*TreeNode, string)
	dfs = func(node *TreeNode, tmp string) {

		tmp = tmp + strconv.Itoa(node.Val)
		if node.Left == nil && node.Right == nil {
			res = append(res, tmp)
			return
		}
		if node.Left != nil {
			dfs(node.Left, tmp+"->")
		}
		if node.Right != nil {
			dfs(node.Right, tmp+"->")
		}
	}

	dfs(root, "")

	return res
}
