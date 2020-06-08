using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCodeTop100CSharp.剑指Offer2 {
    class _0025_合并两个排序的链表 {
        public ListNode MergeTwoLists(ListNode l1, ListNode l2) {
            var head = new ListNode(-1);
            var pre = head;
            while (l1 != null && l2 != null) {
                if (l1.Val < l2.Val) {
                    head.Next = l1;
                    l1 = l1.Next;
                } else {
                    head.Next = l2;
                    l2 = l2.Next;
                }
                head = head.Next;
            }

            if (l1 == null)
                head.Next = l2;
            else
                head.Next = l1;

            return pre.Next;
        }
    }
}
