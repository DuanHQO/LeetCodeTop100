package algirithm

import "leetcodetop100/kit"

type ListNode kit.ListNode

func sortList(head *kit.ListNode) *kit.ListNode {
	if head == nil {
		return nil
	}

	result := mergeSortNode(head)
	return result
}

func mergeSortNode(root *kit.ListNode) *kit.ListNode {
	if root.Next == nil {
		return root
	}

	slow := root
	fast := root

	leftRoot := slow

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	rightRoot := slow.Next
	slow.Next = nil

	left := mergeSortNode(leftRoot)
	right := mergeSortNode(rightRoot)

	result := mergeNode(left, right)
	return result
}

func mergeNode(left, right *kit.ListNode) *kit.ListNode {
	pre := &kit.ListNode{
		Val:  0,
		Next: nil,
	}

	temp := pre

	for left != nil && right != nil {
		if left.Val < right.Val {
			pre.Next = left
			left = left.Next
		} else {
			pre.Next = right
			right = right.Next
		}
		pre = pre.Next
	}

	if left == nil {
		pre.Next = right
	} else {
		pre.Next = left
	}

	return temp.Next
}
