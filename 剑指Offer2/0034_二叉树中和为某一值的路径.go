package 剑指Offer2

func PathSum(root *TreeNode, sum int) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	var helper func(node *TreeNode, left int, tmp []int)
	helper = func(node *TreeNode, left int, tmp []int) {
		if node.Left == nil && node.Right == nil {
			if left == 0 {
				slice := make([]int, len(tmp), len(tmp))
				copy(slice, tmp)
				res = append(res, slice)
			}
			return
		}
		if node.Left != nil {
			tmp = append(tmp, node.Left.Val)
			helper(node.Left, left-node.Left.Val, tmp)
			tmp = tmp[:len(tmp)-1]
		}
		if node.Right != nil {
			tmp = append(tmp, node.Right.Val)
			helper(node.Right, left-node.Right.Val, tmp)
			tmp = tmp[:len(tmp)-1]
		}
	}

	helper(root, sum-root.Val, []int{root.Val})

	return res
}
