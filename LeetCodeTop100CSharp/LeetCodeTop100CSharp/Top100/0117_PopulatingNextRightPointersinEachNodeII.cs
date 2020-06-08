using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCodeTop100CSharp {
    class _0117_PopulatingNextRightPointersinEachNodeII {
        public ListNode Connect(ListNode root) {
            if (root == null) {
                return null;
            }

            var queue = new Queue<ListNode>();
            queue.Enqueue(root);
            while (queue.Count > 0) {
                var len = queue.Count;
                while (len > 0) {
                    var node = queue.Dequeue();
                    if (len > 1) {
                        var next = queue.Peek();
                        node.Next = next;
                    }
                    if (node.Left != null) {
                        queue.Enqueue(node.Left);
                    }
                    if (node.Right != null) {
                        queue.Enqueue(node.Right);
                    }
                    len--;
                }
            }

            return root;
        }
    }
}
