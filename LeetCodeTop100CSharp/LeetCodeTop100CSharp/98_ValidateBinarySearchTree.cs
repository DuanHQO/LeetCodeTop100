using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCodeTop100CSharp {
    class _98_ValidateBinarySearchTree {
        public bool IsValidBST(TreeNode root) {
            var stack = new Stack<TreeNode>();
            var node = root;
            var val = double.NegativeInfinity;
            while (node != null || stack.Count > 0) {
                while (node!= null) {
                    stack.Push(node);
                    node = node.left;
                }
                node = stack.Pop();
                if(val >= node.val) {
                    return false;
                }
                val = node.val;
                node = node.right;
            }

            return true;
        }
    }
}
