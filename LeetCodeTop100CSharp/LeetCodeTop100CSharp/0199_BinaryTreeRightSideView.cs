using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCodeTop100CSharp {
    class _0199_BinaryTreeRightSideView {
        public IList<int> RightSideView(TreeNode root) {
            var res = new Queue<int>();
            if(root == null) {
                return res.ToArray();
            }

            var queue = new Queue<TreeNode>();
            queue.Enqueue(root);
            while (queue.Count > 0) {
                var len = queue.Count;
                while (len > 0) {
                    var node = queue.Dequeue();
                    if(len == 1) {
                        res.Enqueue(node.val);
                    }
                    if(node.left!= null) {
                        queue.Enqueue(node.left);
                    }
                    if(node.right != null) {
                        queue.Enqueue(node.right);
                    }
                    len--;
                }
            }

            return res.ToArray();
        }
    }
}
