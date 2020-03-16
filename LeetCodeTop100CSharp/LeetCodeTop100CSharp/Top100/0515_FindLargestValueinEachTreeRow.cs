using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCodeTop100CSharp {
    class _0515_FindLargestValueinEachTreeRow {
        public IList<int> LargestValues(TreeNode root) {
            var res = new List<int>();
            if (root == null)
                return res.ToArray();

            var queue = new Queue<TreeNode>();
            queue.Enqueue(root);
            while (queue.Count > 0) {
                var len = queue.Count;
                var max = queue.Peek().val;
                while (len > 0) {
                    var node = queue.Dequeue();
                    max = Math.Max(node.val, max);
                    if (node.left != null)
                        queue.Enqueue(node.left);
                    if (node.right != null)
                        queue.Enqueue(node.right);
                    if (len == 1)
                        res.Add(max);
                    len--;
                }
            }


            return res.ToArray();
        }
    }
}
