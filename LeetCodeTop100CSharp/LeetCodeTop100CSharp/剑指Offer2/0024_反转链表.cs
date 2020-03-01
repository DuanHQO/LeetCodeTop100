using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCodeTop100CSharp.剑指Offer2 {
    class _0024_反转链表 {
        public ListNode ReverseList(ListNode head) {
            var res = new List<int>();
            if (head == null)
                return null;

            ListNode pre = null;
            while (head != null) {
                var next = head.next;
                head.next = pre;
                pre = head;
                head = next;
            }

            return pre;

        }
    }
}
