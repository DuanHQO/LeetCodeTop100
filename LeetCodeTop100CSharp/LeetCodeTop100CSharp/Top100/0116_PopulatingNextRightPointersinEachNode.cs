using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCodeTop100CSharp {
    class _0116_PopulatingNextRightPointersinEachListNode {
        public ListNode Connect(ListNode root) {
            if (root == null) {
                return null;
            }

            var queue = new Queue<ListNode>();
            queue.Enqueue(root);
            while (queue.Count > 0) {
                var len = queue.Count;
                while (len > 0) {
                    var ListNode = queue.Dequeue();
                    if(len > 1) {
                        var next = queue.Peek();
                        ListNode.Next = next;
                    }
                    if (ListNode.Left != null) {
                        queue.Enqueue(ListNode.Left);
                    }
                    if (ListNode.Right != null) {
                        queue.Enqueue(ListNode.Right);
                    }
                    len--;
                }
            }

            return root;
        }
    }
}
