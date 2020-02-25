using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCodeTop100CSharp {
    class _0222_CountCompleteTreeNodes {
        public int CountNodes(TreeNode root) {
            if (root == null) {
                return 0;
            }

            var res = 0;
            var result = new Queue<int>();
            var stack = new Stack<TreeNode>();
            stack.Push(root);
            while (stack.Count > 0) {
                var node = stack.Pop();
                res++;
                if (node.right != null)
                    stack.Push(node.right);
                if (node.left != null)
                    stack.Push(node.left);
            }
            return res;
        }
    }
}
