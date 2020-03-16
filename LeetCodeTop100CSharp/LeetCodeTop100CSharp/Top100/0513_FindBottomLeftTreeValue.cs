using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCodeTop100CSharp {
    class _0513_FindBottomLeftTreeValue {
        public int FindBottomLeftValue(TreeNode root) {
            var res = 0;
            if (root == null)
                return res;

            var queue = new Queue<TreeNode>();
            queue.Enqueue(root);
            while (queue.Count > 0) {
                var len = queue.Count;
                var tmp = new Queue<TreeNode>();
                while (len > 0) {
                    var node = queue.Dequeue();
                    tmp.Enqueue(node);
                    if (node.left != null)
                        queue.Enqueue(node.left);
                    if (node.right != null)
                        queue.Enqueue(node.right);
                    len--;
                }
                if (queue.Count == 0)
                    res = tmp.Dequeue().val;
            }

            return res;
        }
    }
}
