package algirithm

func generateTrees(n int) []*TreeNode {
	var res []*TreeNode
	if n == 0 {
		return res
	}

	var helper func(int, int) []*TreeNode
	helper = func(start int, end int) []*TreeNode {
		var allTrees []*TreeNode
		if start > end {
			allTrees = append(allTrees, nil)
			return allTrees
		}

		for i := start; i <= end; i++ {
			leftTrees := helper(start, i-1)
			rightTrees := helper(i+1, end)

			for _, l := range leftTrees {
				for _, r := range rightTrees {
					currentTree := &TreeNode{
						Val:   i,
						Left:  l,
						Right: r,
					}
					allTrees = append(allTrees, currentTree)
				}
			}
		}

		return allTrees
	}

	res = helper(1, n)

	return res
}
