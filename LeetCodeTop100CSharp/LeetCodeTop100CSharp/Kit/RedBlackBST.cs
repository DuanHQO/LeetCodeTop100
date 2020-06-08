using System;

namespace LeetCodeTop100CSharp {
    public class RedBlackBST<TKey, TValue> : BST<TKey, TValue> where TKey: IComparable<TKey> {

        private static readonly bool RED = true;
        private static readonly bool BLACK = false;

        private ListNode root;

        private bool isRed(ListNode h) {
            return false;
        }
        
        private TreeNode<TKey, TValue> rotateLeft(TreeNode<TKey, TValue> h) {
            var x = h.Right;
            h.Right = x.Left;
            x.Left = h;
            x.color = h.color;
            h.color = RED;
            x.N = h.N;
            return x;
        }
    }
}