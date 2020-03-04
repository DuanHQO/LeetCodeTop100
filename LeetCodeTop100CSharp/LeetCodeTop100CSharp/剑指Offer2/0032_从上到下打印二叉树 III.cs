using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCodeTop100CSharp.剑指Offer2 {
    class _0032_从上到下打印二叉树_III {
        public IList<IList<int>> LevelOrder(TreeNode root) {
            var list = new List<int[]>();
            if (root == null)
                return list.ToArray();

            var idx = 1;
            var queue = new Queue<TreeNode>();
            queue.Enqueue(root);
            while (queue.Count > 0) {
                var len = queue.Count;
                var q = new Queue<int>();
                var s = new Stack<int>();
                while (len > 0) {
                    var node = queue.Dequeue();
                    if ((idx & 1) == 1) {
                        q.Enqueue(node.val);
                    } else {
                        s.Push(node.val);
                    }
                    if (node.left != null)
                        queue.Enqueue(node.left);
                    if (node.right != null)
                        queue.Enqueue(node.right);
                    len--;
                }
                if ((idx & 1) == 1) {
                    list.Add(q.ToArray());
                } else {
                    list.Add(s.ToArray());
                }
                idx++;
            }

            return list.ToArray();
        }
    }
}
