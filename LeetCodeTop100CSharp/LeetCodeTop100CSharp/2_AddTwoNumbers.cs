using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCodeTop100CSharp {
    class _2_AddTwoNumbers {
        public static ListNode AddTwoNumbers(ListNode l1, ListNode l2) {

            var head = new ListNode(-1);
            var current = head;

            while (l1 != null && l2 != null) {
                var add = l1.val + l2.val;
                var node = new ListNode(add);
                current.next = node;
                current = node;
                l1 = l1.next;
                l2 = l2.next;
            }

            while (l1 == null && l2 != null) {
                var node = new ListNode(l2.val);
                current.next = node;
                current = node;
                l2 = l2.next;
            }

            while (l2 == null && l1 != null) {
                var node = new ListNode(l2.val);
                current.next = node;
                current = node;
                l1 = l1.next;
            }

            current = head.next;

            while (current != null) {
                if(current.val >= 10) {
                    if(current.next != null) {
                        current.next.val += 1;
                    } else {
                        current.next = new ListNode(1);
                    }
                }
                current.val = current.val % 10;
                current = current.next;
            }

            return head.next;
        }
    }

    public class ListNode {
        public int val;
        public ListNode next;
        public ListNode(int x) { val = x; }
    }
}
