package algirithm

func pathSumII(root *TreeNode, sum int) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	var helper func(int, *TreeNode, []int)
	helper = func(tmpSum int, head *TreeNode, tmpSlice []int) {
		if head == nil {
			return
		}

		if head.Left == nil && head.Right == nil {
			if tmpSum == sum {
				slice := make([]int, len(tmpSlice), len(tmpSlice))
				copy(slice, tmpSlice)
				res = append(res, slice)
				return
			} else {
				return
			}
		}

		if head.Left != nil {
			tmpSlice = append(tmpSlice, head.Left.Val)
			helper(tmpSum+head.Left.Val, head.Left, tmpSlice)
			tmpSlice = tmpSlice[:len(tmpSlice)-1]
		}
		if head.Right != nil {
			tmpSlice = append(tmpSlice, head.Right.Val)
			helper(tmpSum+head.Right.Val, head.Right, tmpSlice)
			tmpSlice = tmpSlice[:len(tmpSlice)-1]
		}
	}

	helper(root.Val, root, []int{root.Val})

	return res
}
