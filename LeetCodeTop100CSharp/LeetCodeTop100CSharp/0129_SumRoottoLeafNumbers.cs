using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCodeTop100CSharp {
    class _0129_SumRoottoLeafNumbers {
        public int SumNumbers(TreeNode root) {
            if (root == null) {
                return 0;
            }

            var res = 0;
            var queue = new Queue<TreeNode>();
            queue.Enqueue(root);
            while (queue.Count > 0) {
                var len = queue.Count;
                while (len > 0) {
                    var node = queue.Dequeue();
                    if(node.left != null) {
                        node.left.val += node.val * 10;
                        queue.Enqueue(node.left);
                    }
                    if (node.right != null) {
                        node.right.val += node.val * 10;
                        queue.Enqueue(node.right);
                    }
                    if(node.left == null && node.right == null) {
                        res += node.val;
                    }
                    len--;
                }
            }

            return res;
        }
    }
}
