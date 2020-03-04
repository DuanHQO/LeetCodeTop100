using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCodeTop100CSharp.剑指Offer2 {
    class _0032_从上到下打印二叉树 {
        public int[] LevelOrder(TreeNode root) {
            var res = new List<int>();
            if (root == null) return res.ToArray();

            var queue = new Queue<TreeNode>();
            queue.Enqueue(root);
            while (queue.Count > 0) {
                var len = queue.Count;
                while (len > 0) {
                    var node = queue.Dequeue();
                    res.Add(node.val);
                    if (node.left != null)
                        queue.Enqueue(node.left);
                    if (node.right != null)
                        queue.Enqueue(node.right);
                    len--;
                }
            }

            return res.ToArray();
        }
    }
}
