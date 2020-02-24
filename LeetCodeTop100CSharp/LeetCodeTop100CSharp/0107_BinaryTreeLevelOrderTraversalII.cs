using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCodeTop100CSharp {
    class _0107_BinaryTreeLevelOrderTraversalII {
        public static IList<IList<int>> LevelOrderBottom(TreeNode root) {
            var res = new Stack<int[]>();
            if(root == null) {
                return res.ToArray();
            }

            var queue = new Queue<TreeNode>();
            queue.Enqueue(root);
            while (queue.Count > 0) {
                var len = queue.Count;
                var tmpList = new List<int>();
                while (len > 0) {
                    var node = queue.Dequeue();
                    tmpList.Add(node.val);
                    if(node.left != null) {
                        queue.Enqueue(node.left);
                    }
                    if(node.right != null) {
                        queue.Enqueue(node.right);
                    }
                    len--;
                }
                res.Push(tmpList.ToArray());
            }

            return res.ToArray();
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
            LevelOrderBottom(node1);
        }
    }
}
