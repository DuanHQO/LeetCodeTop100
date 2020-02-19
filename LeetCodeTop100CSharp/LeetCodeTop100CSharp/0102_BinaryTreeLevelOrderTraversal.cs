using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCodeTop100CSharp {
    class _102_BinaryTreeLevelOrderTraversal {

        public IList<IList<int>> LevelOrder(TreeNode root) {
            if (root == null) return null;

            var result = new List<int[]>();

            var queue = new Queue<TreeNode>();

            queue.Enqueue(root);
            while (queue.Count > 0) {
                var length = queue.Count;
                var element = new int[length];
                for (int i = 0; i < length; i++) {
                    var node = queue.Dequeue();
                    element[i] = node.val;
                    if (node.left != null) {
                        queue.Enqueue(node.left);
                    }
                    if (node.right != null) {
                        queue.Enqueue(root.right);
                    }
                }
                result.Add(element);
            }

            return result.ToArray();
        }
    }
}
