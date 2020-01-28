using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCodeTop100CSharp {
    class _104_MaximumDepthofBinaryTree {
        public static int MaxDepth(TreeNode root) {
            if (root == null) return 0;

            var queue = new Queue<TreeNode>();
            var depth = 0;
            queue.Enqueue(root);
            while (queue.Count > 0) {
                var length = queue.Count;
                for (int i = 0; i < length; i++) {
                    var node = queue.Dequeue();
                    if (node.left != null) {
                        queue.Enqueue(node.left);
                    }
                    if (node.right != null) {
                        queue.Enqueue(node.right);
                    }
                }
                depth ++;
            }
            return depth;
        }
    }
}
