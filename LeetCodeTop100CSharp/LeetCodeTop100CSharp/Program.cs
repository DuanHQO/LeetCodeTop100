using System;

namespace LeetCodeTop100CSharp {
    class Program {
        static void Main(string[] args) {

            var node1 = new ListNode(9);
            var node2 = new ListNode(1);
            var node3 = new ListNode(9);
            //node1.next = node2;
            //node2.next = node3;

            var node4 = new ListNode(9);
            var node5 = new ListNode(9);
            var node6 = new ListNode(9);
            var node7 = new ListNode(9);
            var node8 = new ListNode(9);
            var node9 = new ListNode(9);
            var node10 = new ListNode(9);
            var node11 = new ListNode(9);
            node2.next = node3;
            node3.next = node4;
            node4.next = node5;
            node5.next = node6;
            node6.next = node7;
            node7.next = node8;
            node8.next = node9;
            node9.next = node10;
            node10.next = node11;

            var cur = _2_AddTwoNumbers.AddTwoNumbers(node1, node2);
            while (cur != null) {
                Console.WriteLine("{0} -> ", cur.val);
                cur = cur.next;
            }

            Console.Read();
        }
    }
}
