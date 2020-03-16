using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCodeTop100CSharp {
    class _0103_BinaryTreeZigzagLevelOrderTraversal {
        public IList<IList<int>> ZigzagLevelOrder(TreeNode root) {
            var res = new List<int[]>();
            if (root == null) {
                return res.ToArray();
            }
            var queue = new Queue<TreeNode>();
            queue.Enqueue(root);
            var count = 0;
            while (queue.Count > 0) {
                count++;
                var len = queue.Count;
                var tmpList = new List<int>();
                var tmpStack = new Stack<int>();
                while (len > 0) {
                    var node = queue.Dequeue();
                    if ((count & 1) == 1) {
                        tmpList.Add(node.val);
                    } else {
                        tmpStack.Push(node.val);
                    }
                    if (node.left != null) {
                        queue.Enqueue(node.left);
                    }
                    if (node.right != null) {
                        queue.Enqueue(node.right);
                    }
                    len--;
                }
                if ((count & 1) == 1) {
                    res.Add(tmpList.ToArray());
                } else {
                    res.Add(tmpStack.ToArray());
                }
            }

            return res.ToArray();
        }
    }
}
