using System;

namespace LeetCodeTop100CSharp {

    
    public class TreeNode<TKey, TValue> where TKey : IComparable<TKey> {
        public TKey Key;
        public TValue Value;
        public TreeNode<TKey, TValue> Left;
        public TreeNode<TKey, TValue> Right;
        public int N;
        public bool color;

        public TreeNode(TKey key, TValue value, int N, bool color) {
            this.Key = key;
            this.Value = value;
            this.N = N;
            this.color = color;
        }

        public bool isRed(TreeNode<TKey, TValue> x) {
            if (x == null) return false;
            return x.color == true;
        }

    }

    
    public class TreeNode {
        // public 
        public int val;
        public TreeNode left;
        public TreeNode right;
        public TreeNode(int x) { val = x; }
    }
}