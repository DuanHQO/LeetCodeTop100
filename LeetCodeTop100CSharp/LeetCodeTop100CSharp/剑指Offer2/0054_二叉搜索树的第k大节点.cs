using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCodeTop100CSharp.剑指Offer2 {
    class _0054_二叉搜索树的第k大节点 {
        public int KthLargest(TreeNode root, int k) {
            var res = new List<int>();
            var stack = new Stack<TreeNode>();
            var node = root;
            while (node != null || stack.Count > 0) {
                while (node != null) {
                    stack.Push(node);
                    node = node.left;
                }
                node = stack.Pop();
                res.Add(node.val);
                node = node.right;
            }
            return res[res.Count - k];
        }
    }
}
