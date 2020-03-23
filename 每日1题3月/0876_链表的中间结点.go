package 每日1题3月

func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	//if head.Next == nil {
	//	return head
	//}

	slow := head
	fast := head.Next

	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	if fast == nil {
		return slow
	} else {
		return slow.Next
	}

}
