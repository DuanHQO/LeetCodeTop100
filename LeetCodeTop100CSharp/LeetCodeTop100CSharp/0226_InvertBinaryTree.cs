using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCodeTop100CSharp {
    class _226_InvertBinaryTree {

        public static TreeNode InvertTree(TreeNode root) {
            if (root == null) return null;
            var stack = new Stack<TreeNode>();
            stack.Push(root);
            while (stack.Count > 0) {
                var node = stack.Pop();
                if(node.right != null) {
                    stack.Push(node.right);
                }
                if(node.left != null) {
                    stack.Push(node.left);
                }
                var left = node.left;
                node.left = node.right;
                node.right = left;
            }
            return root;
        }
    }
}
