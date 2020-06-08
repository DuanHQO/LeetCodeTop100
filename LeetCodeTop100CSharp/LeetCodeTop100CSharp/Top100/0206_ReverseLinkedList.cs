using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCodeTop100CSharp {
    class _206_ReverseLinkedList {
        public static ListNode ReverseList(ListNode head) {
            if (head == null) return null;

            ListNode pre = null;
            ListNode next = null;
            while (head!= null) {
                next = head.Next;
                head.Next = pre;
                pre = head;
                head = next;
            }
            return pre;
        }
    }
}
