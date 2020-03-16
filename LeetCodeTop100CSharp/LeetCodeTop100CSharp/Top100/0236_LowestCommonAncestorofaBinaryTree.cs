using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCodeTop100CSharp {
    class _0236_LowestCommonAncestorofaBinaryTree {
        public TreeNode LowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {
            var preOrderList = new List<TreeNode>();
            var inOrderList = new List<TreeNode>();

            var stack = new Stack<TreeNode>();
            stack.Push(root);
            while (stack.Count > 0) {
                var nodeTmp = stack.Pop();
                preOrderList.Add(nodeTmp);
                if (nodeTmp.right != null)
                    stack.Push(nodeTmp.right);
                if (nodeTmp.left != null)
                    stack.Push(nodeTmp.left);
            }

            var node = root;
            while (node != null || stack.Count > 0) {
                while (node != null) {
                    stack.Push(node);
                    node = node.left;
                }
                node = stack.Pop();
                inOrderList.Add(node);
                node = node.right;
            }

            var preOrderArray = preOrderList.ToArray();
            var inOrderArray = inOrderList.ToArray();

            var preOrderMap = new Dictionary<TreeNode, int>();
            var inOrderMap = new Dictionary<TreeNode, int>();

            for (int i = 0; i < preOrderArray.Length; i++) {
                preOrderMap.Add(preOrderArray[i], i);
                inOrderMap.Add(inOrderArray[i], i);
            }

            var lo = Math.Min(inOrderMap[p], inOrderMap[q]);
            var hi = Math.Max(inOrderMap[p], inOrderMap[q]);

            Console.WriteLine("lo {0} hi {1}", lo, hi);

            for (int i = 0; i <= preOrderMap[q]; i++) {
                var inordedIdx = inOrderMap[preOrderArray[i]];
                Console.WriteLine("idx {0}", inordedIdx);
                if (lo <= inordedIdx && inordedIdx <= hi) {
                    return preOrderArray[i];
                }
            }

            return null;
        }
    }
}
