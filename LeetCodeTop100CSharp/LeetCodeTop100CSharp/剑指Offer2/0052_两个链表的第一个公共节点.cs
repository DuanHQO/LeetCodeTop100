using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCodeTop100CSharp.剑指Offer2 {
    class _0052_两个链表的第一个公共节点 {
        public ListNode GetIntersectionNode(ListNode headA, ListNode headB) {
            if (headA == null || headB == null)
                return null;

            var map = new Dictionary<ListNode, ListNode>();
            ListNode pre = new ListNode(-1);
            ListNode back = new ListNode(-1);
            pre.next = headA;
            while (headA != null) {
                if (!map.ContainsKey(headA)) {
                    map.Add(headA, pre);
                }
                pre = headA;
                headA = headA.next;
            }
            while (headB != null) {
                if (!map.ContainsKey(headB)) {
                    map.Add(headB, pre);
                } else {
                    if (map[headB] != back) {
                        return headB;
                    }
                }
                back = headB;
                headB = headB.next;
            }
            return null;
        }
    }
}
