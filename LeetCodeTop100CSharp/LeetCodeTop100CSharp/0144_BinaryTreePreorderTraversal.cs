using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCodeTop100CSharp {
    class _0144_BinaryTreePreorderTraversal {
        public IList<int> PreorderTraversal(TreeNode root) {
            var res = new List<int>();
            if(root == null) {
                return res.ToArray();
            }

            var stack = new Stack<TreeNode>();
            stack.Push(root);
            while (stack.Count > 0) {
                var node = stack.Pop();
                res.Add(node.val);
                if (node.right != null) {
                    stack.Push(node.right);
                }
                if (node.left != null) {
                    stack.Push(node.left);
                }
            }

            return res.ToArray();
        }
    }
}
