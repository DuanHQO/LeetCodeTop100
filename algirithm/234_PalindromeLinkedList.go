package algirithm

import "leetcodetop100/kit"

func isPalindrome(head *ListNode) bool {

	pre := head

	stack := kit.NewStack()
	for head != nil {
		stack.Push(head.Val)
		head = head.Next
	}

	head = pre

	for head != nil {
		val := stack.Pop()
		if val != head.Val {
			return false
		}
		head = head.Next
	}

	return true
}
