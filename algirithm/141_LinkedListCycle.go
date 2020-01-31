package algirithm

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	//双指针，如果是环形的，快指针总能追上慢指针
	slow := head
	fast := slow.Next

	for slow != nil && fast != nil && fast.Next != nil && slow != fast {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow == fast
}
