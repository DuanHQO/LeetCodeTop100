using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCodeTop100CSharp {
    class _617_MergeTwoBinaryTrees {

        public static TreeNode MergeTrees(TreeNode t1, TreeNode t2) {

            var stack1 = new Stack<TreeNode>();
            var stack2 = new Stack<TreeNode>();

            stack1.Push(t1);
            stack2.Push(t2);
            if(t1 == null) {
                return t2;
            }
            if(t2 == null) {
                return t1;
            }
            while (stack1.Count > 0) {
                var node1 = stack1.Pop();
                var node2 = stack2.Pop();

                node1.val += node2.val;

                if(node1.right != null && node2.right != null) {
                    stack1.Push(node1.right);
                    stack2.Push(node2.right);
                } else if(node2.right != null) {
                    node1.right = node2.right;
                }
                if(node1.left != null && node2.left != null) {
                    stack1.Push(node1.left);
                    stack2.Push(node2.left);
                } else if(node2.left != null) {
                    node1.left = node2.left;
                }
            }

            return t1;
        }

    }

    public class TreeNode {
        public int val;
        public TreeNode left;
        public TreeNode right;
        public TreeNode(int x) { val = x; }
    }
}
