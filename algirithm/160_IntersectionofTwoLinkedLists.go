package algirithm

import "leetcodetop100/kit"

func getIntersectionNode(headA, headB *kit.ListNode) *kit.ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	nodeA := headA
	nodeB := headB
	for nodeA != nodeB {
		if nodeA == nil {
			nodeA = headB
		} else {
			nodeA = nodeA.Next
		}
		if nodeB == nil {
			nodeB = headA
		} else {
			nodeB = nodeB.Next
		}
	}

	return nodeA
}
