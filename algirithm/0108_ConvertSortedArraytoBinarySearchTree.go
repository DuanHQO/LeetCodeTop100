package algirithm

func sortedArrayToBST(nums []int) *TreeNode {
	if nums == nil {
		return nil
	}

	var helper func(int, int) *TreeNode
	helper = func(lo int, hi int) *TreeNode {
		if lo > hi {
			return nil
		}

		mid := lo + (hi-lo)/2
		root := &TreeNode{Val: nums[mid]}

		root.Left = helper(lo, mid-1)
		root.Right = helper(mid+1, hi)

		return root
	}

	root := helper(0, len(nums)-1)
	return root
}
