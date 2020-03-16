using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCodeTop100CSharp {
    class _0235_LowestCommonAncestorofaBinarySearchTree {
        public TreeNode LowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {
            var pVal = p.val;
            var qVal = q.val;
            var node = root;
            while (node != null) {
                if(pVal < node.val && qVal < node.val) {
                    node = node.left;
                }else if(pVal > node.val && qVal > node.val) {
                    node = node.right;
                } else {
                    return node;
                }
            }
            return null;
        }
    }
}
