using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCodeTop100CSharp.剑指Offer2 {
    class _0028_对称的二叉树 {
        public bool IsSymmetric(TreeNode root) {
            if (root == null) return true;

            var queue = new Queue<TreeNode>();
            queue.Enqueue(root.left);
            queue.Enqueue(root.right);

            while (queue.Count > 0) {
                var left = queue.Dequeue();
                var right = queue.Dequeue();

                if (left == null && right == null)
                    continue;
                if (left == null || right == null)
                    return false;

                if (left.val != right.val)
                    return false;

                queue.Enqueue(left.left);
                queue.Enqueue(right.right);
                queue.Enqueue(left.right);
                queue.Enqueue(right.left);
            }

            return true;
        }

        
    }
}
