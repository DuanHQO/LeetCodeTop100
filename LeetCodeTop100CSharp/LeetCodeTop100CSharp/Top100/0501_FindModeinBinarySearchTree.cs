using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCodeTop100CSharp {
    class _0501_FindModeinBinarySearchTree {
        List<int> res = new List<int>();
        int maxTimes = 0;
        int thisTimes = 0;
        TreeNode pre = null;
        public int[] FindMode(TreeNode root) {
            Inorder(root);
            return res.ToArray();
        }

        public void Inorder(TreeNode node) {
            if (node == null)
                return;
            Inorder(node.left);
            if (pre != null && node.val == pre.val)
                thisTimes++;
            else
                thisTimes = 1;
            if (thisTimes == maxTimes) {
                res.Add(node.val);
            } else if (thisTimes > maxTimes) {
                maxTimes = thisTimes;
                res.Clear();
                res.Add(node.val);
            }
            pre = node;
            Inorder(node.right);
        }
    }
}
