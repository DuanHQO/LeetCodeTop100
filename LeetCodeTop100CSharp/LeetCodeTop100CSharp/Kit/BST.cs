using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCodeTop100CSharp {
    public class BST<TKey, TValue> where TKey:IComparable<TKey> {
        private TreeNode<TKey, TValue> root;

        public int Size() {
            return size(root);
        }

        private int size(TreeNode<TKey, TValue> root) {
            if (root == null) return 0;
            else return root.N;
        }

        public TValue Get(TKey key) {
            return get(root, key);
        }

        private TValue get(TreeNode<TKey, TValue> root, TKey key) {
            if (root == null) return default(TValue);
            var cmp = key.CompareTo(root.Key);
            if (cmp < 0) return get(root.Left, key);
            else if (cmp > 0) return get(root.Right, key);
            else return root.Value;
        }


        public void Put(TKey key, TValue value) {
            root = put(root, key, value);
        }

        private TreeNode<TKey, TValue> put(TreeNode<TKey, TValue> root, TKey key, TValue value) {
            if (root == null) return new TreeNode<TKey, TValue>(key, value, 1, false);
            var cmp = key.CompareTo(root.Key);
            if (cmp < 0) root.Left = put(root.Left, key, value);
            else if (cmp > 0) root.Right = put(root.Right, key, value);
            else root.Value = value;
            root.N = size(root.Left) + size(root.Right) + 1;
            return root;
        }

        public TKey Min() {
            return min(root).Key;
        }

        private TreeNode<TKey, TValue> min(TreeNode<TKey, TValue> root) {
            if (root.Left == null) return root;
            return min(root.Left);
        }

        public TKey Max() {
            return max(root).Key;
        }

        private TreeNode<TKey, TValue> max(TreeNode<TKey, TValue> root) {
            if (root == null) return root;
            return max(root.Right);
        }

        public TKey Floor(TKey key) {
            var x = floor(root, key);
            if (x == null) return default(TKey);
            return x.Key;
        }

        private TreeNode<TKey, TValue> floor(TreeNode<TKey, TValue> root, TKey key) {
            if (root == null) return null;
            var cmp = key.CompareTo(root.Key);
            if (cmp == 0) return root;
            else if (cmp < 0) return floor(root.Left, key);
            var t = floor(root.Right, key);
            if (t != null) return t;
            else return root;
        }

        public TKey Ceiling(TKey key) {
            return ceiling(root, key).Key;
        }

        private TreeNode<TKey, TValue> ceiling(TreeNode<TKey, TValue> root, TKey key) {
            if (root == null) return null;
            var cmp = key.CompareTo(root.Key);
            if (cmp == 0) return root;
            else if (cmp > 0) return ceiling(root.Right, key);
            var t = ceiling(root.Left, key);
            if (t != null) return t;
            return root;
        }

        public TKey Select(int k) {
            return select(root, k).Key;
        }

        private TreeNode<TKey, TValue> select(TreeNode<TKey, TValue> root, int k) {
            if (root == null) return null;
            var t = size(root.Left);
            if (t > k) return select(root.Left, k);
            else if (t < k) return select(root.Right, k - t - 1);
            else return root;
        }

        public int Rank(TKey key) {
            return rank(key, root);
        }

        private int rank(TKey key, TreeNode<TKey, TValue> root) {
            if (root == null) return 0;
            var cmp = key.CompareTo(root.Key);
            if (cmp < 0) return rank(key, root.Left);
            else if (cmp > 0) return 1 + size(root) + rank(key, root.Right);
            else return root.N;
        }

        public void DeleteMin() {
            root = deleteMin(root);
        }

        private TreeNode<TKey, TValue> deleteMin(TreeNode<TKey, TValue> root) {
            if (root.Left == null) return root.Right;
            root.Left = deleteMin(root.Left);
            root.N = size(root.Left) + size(root.Right) + 1;
            return root;
        }

        public void Delete(TKey key) {
            root = delete(root, key);
        }

        private TreeNode<TKey, TValue> delete(TreeNode<TKey, TValue> root, TKey key) {
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

        private IEnumerable<TKey> keys(TKey lo, TKey hi) {
            var queue = new Queue<TKey>();
            keys(root, queue, lo, hi);
            return queue;
        }

        private void keys(TreeNode<TKey, TValue> root, Queue<TKey> queue, TKey lo, TKey hi) {
            if (root == null) return;
            var cmplo = lo.CompareTo(root.Key);
            var cmphi = hi.CompareTo(root.Key);
            if (cmplo < 0) keys(root.Left, queue, lo, hi);
            if (cmplo <= 0 && cmphi >= 0) queue.Enqueue(root.Key);
            if (cmphi > 0) keys(root.Right, queue, lo, hi);
        }
    }
}
