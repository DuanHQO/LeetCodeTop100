namespace LeetCodeTop100CSharp {
    class _21_MergeTwoSortedLists {
        public ListNode MergeTwoLists(ListNode l1, ListNode l2) {
            if(l1 == null) {
                return l2;
            }
            if(l2 == null) {
                return l1;
            }
            ListNode cur = new ListNode(-1);
            ListNode pre = cur;
            while (l1 != null && l2 != null) {
                if(l1.Val < l2.Val) {
                    cur.Next = l1;
                    l1 = l1.Next;
                } else {
                    cur.Next = l2;
                    l2 = l2.Next;
                }
                cur = cur.Next;
            }

            if(l1 == null) {
                cur.Next = l2;
            } else {
                cur.Next = l1;
            }
            return pre.Next;
        }
    }
}
