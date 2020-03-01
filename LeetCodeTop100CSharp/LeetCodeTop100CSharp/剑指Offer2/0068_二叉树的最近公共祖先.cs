using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCodeTop100CSharp.剑指Offer2 {
    class _0068_二叉树的最近公共祖先 {
        private TreeNode ans;
        public TreeNode LowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {
            RecurseTree(root, p, q);
            return ans;
        }

        private bool RecurseTree(TreeNode node, TreeNode p, TreeNode q) {
            if (node == null) {
                return false;
            }

            var left = RecurseTree(node.left, p , q) ? 1 : 0;
            var right = RecurseTree(node.right, p, q) ? 1 : 0;
            var mid = (node == p || node == q) ? 1 : 0;

            if (mid + left + right >= 2) {
                ans = node;
            }

            return (mid + left + right > 0);
        }
    }
}
