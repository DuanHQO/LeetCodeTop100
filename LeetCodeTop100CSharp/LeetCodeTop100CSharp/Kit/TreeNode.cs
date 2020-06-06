using System;

namespace LeetCodeTop100CSharp {
    public class TreeNode<TKey, TValue> where TKey : IComparable<TKey> {
        public TKey Key;
        public TValue Value;
        public TreeNode<TKey, TValue> Left;
        public TreeNode<TKey, TValue> Right;
        public TreeNode(TValue x) { Value = x; }
        
        // public bool isRed()
    }
    
    public class TreeNode {
        // public 
        public int val;
        public TreeNode left;
        public TreeNode right;
        public TreeNode(int x) { val = x; }
    }
}