package 剑指Offer2

func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	pre := &ListNode{Val: -1, Next: head}
	back := pre
	for head != nil {
		if head.Val == val {
			pre.Next = head.Next
			break
		}
		pre = head
		head = head.Next
	}

	return back.Next
}
