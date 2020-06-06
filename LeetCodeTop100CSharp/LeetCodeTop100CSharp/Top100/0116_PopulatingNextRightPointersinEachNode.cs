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
                        ListNode.next = next;
                    }
                    if (ListNode.left != null) {
                        queue.Enqueue(ListNode.left);
                    }
                    if (ListNode.right != null) {
                        queue.Enqueue(ListNode.right);
                    }
                    len--;
                }
            }

            return root;
        }
    }
}
