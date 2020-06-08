using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCodeTop100CSharp.剑指Offer2 {
    class _0006_从尾到头打印链表 {
        public int[] ReversePrint(ListNode head) {
            var res = new List<int>();
            if (head == null)
                return res.ToArray();

            var back = head;
            ListNode pre = null;
            while (back != null) {
                var next = back.Next;
                back.Next = pre;
                pre = back;
                back = next;
            }

            while (pre != null) {
                res.Add(pre.Val);
                pre = pre.Next;
            }

            return res.ToArray();
        }
    }
}
