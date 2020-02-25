using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCodeTop100CSharp {
    class _0404_SumofLeftLeaves {
        public int SumOfLeftLeaves(TreeNode root) {
            var node = root;
            var stack = new Stack<TreeNode>();
            var res = 0;
            while (node != null || stack.Count > 0) {
                while (node != null) {
                    stack.Push(node);
                    node = node.left;
                }
                node = stack.Pop();
                if (node.left != null && node.left.left == null && node.left.right == null)
                    res += node.left.val;
                node = node.right;
            }

            return res;
        }
    }
}
