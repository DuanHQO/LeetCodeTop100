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
                        node.next = next;
                    }
                    if (node.left != null) {
                        queue.Enqueue(node.left);
                    }
                    if (node.right != null) {
                        queue.Enqueue(node.right);
                    }
                    len--;
                }
            }

            return root;
        }
    }
}
