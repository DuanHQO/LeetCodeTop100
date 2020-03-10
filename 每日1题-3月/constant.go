package 每日1题_3月

import "leetcodetop100/kit"

type TreeNode = kit.TreeNode
type ListNode = kit.ListNode

type chars []byte

func (v chars) Len() int {
	return len(v)
}

func (v chars) Less(i, j int) bool {
	return v[i] < v[j]
}

func (v chars) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}
