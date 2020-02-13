package algirithm

import "leetcodetop100/kit"

func removeNthFromEnd(head *kit.ListNode, n int) *kit.ListNode {
	if head == nil {
		return nil
	}

	pre := &kit.ListNode{
		Val:  0,
		Next: head,
	}

	slow := pre
	fast := pre

	for i := 1; i <= n+1; i++ {
		fast = fast.Next
	}

	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}

	slow.Next = slow.Next.Next

	return pre.Next
}
