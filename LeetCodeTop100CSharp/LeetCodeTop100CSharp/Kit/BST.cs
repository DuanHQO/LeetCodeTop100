using System;
using System.Collections.Generic;

namespace LeetCodeTop100CSharp {
    public class BST<TKey, TValue> where TKey:IComparable<TKey> {
        private TreeNode<TKey, TValue> _root;

        public int Size() {
            return size(_root);
        }

        protected int size(TreeNode<TKey, TValue> root) {
            if (root == null) return 0;
            return root.N;
        }

        public TValue Get(TKey key) {
            return get(_root, key);
        }

        protected TValue get(TreeNode<TKey, TValue> x, TKey key) {
            if (x == null) return default(TValue);
            var cmp = key.CompareTo(x.Key);
            if (cmp < 0) return get(x.Left, key);
            return cmp > 0 ? get(x.Right, key) : x.Value;
        }


        public void Put(TKey key, TValue value) {
            _root = put(_root, key, value);
        }

        protected TreeNode<TKey, TValue> put(TreeNode<TKey, TValue> root, TKey key, TValue value) {
            if (root == null) return new TreeNode<TKey, TValue>(key, value, 1, false);
            var cmp = key.CompareTo(root.Key);
            if (cmp < 0) root.Left = put(root.Left, key, value);
            else if (cmp > 0) root.Right = put(root.Right, key, value);
            else root.Value = value;
            root.N = size(root.Left) + size(root.Right) + 1;
            return root;
        }

        public TKey Min() {
            return min(_root).Key;
        }

        protected TreeNode<TKey, TValue> min(TreeNode<TKey, TValue> root) {
            if (root.Left == null) return root;
            return min(root.Left);
        }

        public TKey Max() {
            return max(_root).Key;
        }

        protected TreeNode<TKey, TValue> max(TreeNode<TKey, TValue> root) {
            if (root == null) return root;
            return max(root.Right);
        }

        public TKey Floor(TKey key) {
            var x = floor(_root, key);
            if (x == null) return default(TKey);
            return x.Key;
        }

        protected TreeNode<TKey, TValue> floor(TreeNode<TKey, TValue> root, TKey key) {
            if (root == null) return null;
            var cmp = key.CompareTo(root.Key);
            if (cmp == 0) return root;
            if (cmp < 0) return floor(root.Left, key);
            var t = floor(root.Right, key);
            return t ?? root;
        }

        public TKey Ceiling(TKey key) {
            return ceiling(_root, key).Key;
        }

        protected TreeNode<TKey, TValue> ceiling(TreeNode<TKey, TValue> root, TKey key) {
            if (root == null) return null;
            var cmp = key.CompareTo(root.Key);
            if (cmp == 0) return root;
            if (cmp > 0) return ceiling(root.Right, key);
            var t = ceiling(root.Left, key);
            return t ?? root;
        }

        public TKey Select(int k) {
            return select(_root, k).Key;
        }

        protected TreeNode<TKey, TValue> select(TreeNode<TKey, TValue> root, int k) {
            if (root == null) return null;
            var t = size(root.Left);
            if (t > k) return select(root.Left, k);
            return t < k ? @select(root.Right, k - t - 1) : root;
        }

        public int Rank(TKey key) {
            return rank(key, _root);
        }

        protected int rank(TKey key, TreeNode<TKey, TValue> root) {
            if (root == null) return 0;
            var cmp = key.CompareTo(root.Key);
            if (cmp < 0) return rank(key, root.Left);
            if (cmp > 0) return 1 + size(root) + rank(key, root.Right);
            return root.N;
        }

        public void DeleteMin() {
            _root = deleteMin(_root);
        }

        protected TreeNode<TKey, TValue> deleteMin(TreeNode<TKey, TValue> root) {
            if (root.Left == null) return root.Right;
            root.Left = deleteMin(root.Left);
            root.N = size(root.Left) + size(root.Right) + 1;
            return root;
        }

        public void Delete(TKey key) {
            _root = delete(_root, key);
        }

        protected TreeNode<TKey, TValue> delete(TreeNode<TKey, TValue> root, TKey key) {
            if (root == null) return null;
            var cmp = key.CompareTo(root.Key);
            if (cmp < 0) root.Left = delete(root.Left, key);
            else if (cmp > 0) root.Right = delete(root.Right, key);
            else {
                if (root.Right == null) return root.Left;
                if (root.Left == null) return root.Right;
                var t = root;
                root = min(root.Right);
                root.Right = deleteMin(t.Right);
                root.Left = t.Left;
            }
            root.N = size(root.Left) + 1 + size(root.Right);
            return root;
        }

        public IEnumerable<TKey> Keys() {
            return keys(Min(), Max());
        }

        protected IEnumerable<TKey> keys(TKey lo, TKey hi) {
            var queue = new Queue<TKey>();
            keys(_root, queue, lo, hi);
            return queue;
        }

        protected void keys(TreeNode<TKey, TValue> root, Queue<TKey> queue, TKey lo, TKey hi) {
            if (root == null) return;
            var cmplo = lo.CompareTo(root.Key);
            var cmphi = hi.CompareTo(root.Key);
            if (cmplo < 0) keys(root.Left, queue, lo, hi);
            if (cmplo <= 0 && cmphi >= 0) queue.Enqueue(root.Key);
            if (cmphi > 0) keys(root.Right, queue, lo, hi);
        }
    }
}
