using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCodeTop100CSharp.剑指Offer2 {
    class _0036_二叉搜索树与双向链表 {
        public ListNode TreeToDoublyList(ListNode root) {
            if (root == null)
                return null;

            var stack = new Stack<ListNode>();
            var node = root;
            ListNode pre = null;

            while (node != null || stack.Count > 0) {
                while (node != null) {
                    stack.Push(node);
                    node = node.left;
                }
                node = stack.Pop();
                if (pre != null) {
                    node.left = pre;
                    pre.right = node;
                }
                pre = node;
                node = node.right;
            }
            pre.right = root;
            return root;
        }
    }
}
