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
                var add = l1.Val + l2.Val;
                var node = new ListNode(add);
                current.Next = node;
                current = node;
                l1 = l1.Next;
                l2 = l2.Next;
            }

            while (l1 == null && l2 != null) {
                var node = new ListNode(l2.Val);
                current.Next = node;
                current = node;
                l2 = l2.Next;
            }

            while (l2 == null && l1 != null) {
                var node = new ListNode(l2.Val);
                current.Next = node;
                current = node;
                l1 = l1.Next;
            }

            current = head.Next;

            while (current != null) {
                if(current.Val >= 10) {
                    if(current.Next != null) {
                        current.Next.Val += 1;
                    } else {
                        current.Next = new ListNode(1);
                    }
                }
                current.Val = current.Val % 10;
                current = current.Next;
            }

            return head.Next;
        }
    }

    // public class ListNode {
    //     public int val;
    //     public ListNode next;
    //     public ListNode(int x) { val = x; }
    // }
}
