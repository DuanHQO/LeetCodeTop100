using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCodeTop100CSharp {
    class _0173_BinarySearchTreeIterator {
        
    }

    public class BSTIterator {

        Queue<int> queue;

        public BSTIterator(TreeNode root) {
            queue = new Queue<int>();
            var stack = new Stack<TreeNode>();
            var node = root;
            while (node != null || stack.Count > 0) {
                while (node != null) {
                    stack.Push(node);
                    node = node.left;
                }
                node = stack.Pop();
                queue.Enqueue(node.val);
                node = node.right;
            }
        }

        /** @return the next smallest number */
        public int Next() {
            return queue.Dequeue();
        }

        /** @return whether we have a next smallest number */
        public bool HasNext() {
            return queue.Count > 0;
        }
    }
}
