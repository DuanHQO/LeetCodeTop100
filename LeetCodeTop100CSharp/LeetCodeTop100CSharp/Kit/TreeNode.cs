using System;

namespace LeetCodeTop100CSharp {
    public class TreeNode<TKey, TValue> where TKey : IComparable<TKey> {
        // public 
        public int val;
        public TreeNode<TKey, TValue> left;
        public TreeNode<TKey, TValue> right;
        public TreeNode(int x) { val = x; }
    }
    
    public class TreeNode {
        // public 
        public int val;
        public TreeNode left;
        public TreeNode right;
        public TreeNode(int x) { val = x; }
    }
}