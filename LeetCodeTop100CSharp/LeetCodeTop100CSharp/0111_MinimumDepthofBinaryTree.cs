using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCodeTop100CSharp {
    class _0111_MinimumDepthofBinaryTree {
        public static int MinDepth(TreeNode root) {
            if (root == null) {
                return 0;
            }

            var queue = new Queue<TreeNode>();
            var height = 0;
            queue.Enqueue(root);
            while (queue.Count > 0) {
                var len = queue.Count;
                height++;
                while (len > 0) {
                    var node = queue.Dequeue();
                    if (node.left == null && node.right == null) {
                        return height;
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
            return height;
        }

        public static void Main(string[] args) {
            var node1 = new TreeNode(3);
            var node2 = new TreeNode(9);
            var node3 = new TreeNode(20);
            node1.left = node2;
            node1.right = node3;
            var node4 = new TreeNode(15);
            var node5 = new TreeNode(7);
            node3.left = node4;
            node3.right = node5;
            MinDepth(node1);
        }
    }
}
