using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCodeTop100CSharp {
    class _0230_KthSmallestElementinaBST {
        public int KthSmallest(TreeNode root, int k) {
            var node = root;
            var stack = new Stack<TreeNode>();
            var list = new List<TreeNode>();
            while (node != null || stack.Count > 0) {
                while (node != null) {
                    stack.Push(node);
                    node = node.left;
                }
                node = stack.Pop();
                list.Add(node);
                node = node.right;
            }
            return list[k - 1].val;
        }
    }
}
