using System;
using System.Collections.Generic;

namespace LeetCodeTop100CSharp {
    public class RedBlackBST<TKey, TValue> where TKey: IComparable<TKey> {

        private static readonly bool RED = true;
        private static readonly bool BLACK = false;

        private TreeNode<TKey, TValue> _root;

        private int Size(TreeNode<TKey, TValue> x) {
            if (x == null) return 0;
            return x.N;
        }

        private bool IsRed(TreeNode<TKey, TValue> x) {
            if (x == null) return false;
            return x.color == RED;
        }

        private void FlipColors(TreeNode<TKey, TValue> x) {
            x.color = RED;
            x.Left.color = BLACK;
            x.Right.color = BLACK;
        }

        private void FlipColorsDesc(TreeNode<TKey, TValue> x) {
            x.color = BLACK;
            x.Left.color = RED;
            x.Right.color = RED;
        }

        public void Put(TKey key, TValue value) {
            _root = Put(_root, key, value);
            _root.color = BLACK;
        }

        public TreeNode<TKey, TValue> Balance(TreeNode<TKey, TValue> x) {
            if (IsRed(x.Right) && !IsRed(x.Left)) x = RotateLeft(x);
            if (IsRed(x.Left) && IsRed(x.Left.Left)) x = RotateRight(x);
            if (IsRed(x.Left) && IsRed(x.Right)) FlipColors(x); 
            
            x.N = Size(x.Left) + 1 + Size(x.Right);
            return x;
        }

        private TreeNode<TKey, TValue> Put(TreeNode<TKey, TValue> x, TKey key, TValue value) {
            if (x == null) return new TreeNode<TKey, TValue>(key, value, 1, RED);
            var cmp = key.CompareTo(x.Key);
            if (cmp < 0) x.Left = Put(x.Left, key, value);
            if (cmp > 0) x.Right = Put(x.Right, key, value);
            x.Value = value;

            return Balance(x);
        }

        public TValue Get(TKey key) {
            return Get(_root, key);
        }

        public TValue Get(TreeNode<TKey, TValue> x, TKey key) {
            while (true) {
                if (x == null) return default(TValue);
                var cmp = key.CompareTo(x.Key);
                if (cmp < 0) {
                    x = x.Left;
                    continue;
                }
                if (cmp > 0) {
                    x = x.Right;
                    continue;
                }
                return x.Value;
            }
        }

        public TKey Min() {
            return Min(_root).Key;
        }

        private TreeNode<TKey, TValue> Min(TreeNode<TKey, TValue> x) {
            if (x.Left == null) return x;
            return Min(x.Left);
        }

        public TKey Max() {
            return Max(_root).Key;
        }

        private TreeNode<TKey, TValue> Max(TreeNode<TKey, TValue> x) {
            if (x.Right == null) return x;
            return Max(x.Right);
        }

        public TKey Floor(TKey key) {
            var x = Floor(_root, key);
            return x == null ? default(TKey) : x.Key;
        }

        private TreeNode<TKey, TValue> Floor(TreeNode<TKey, TValue> x, TKey key) {
            if (x == null) return null;
            var cmp = key.CompareTo(x.Key);
            if (cmp == 0) return x;
            if (cmp < 0) return Floor(x.Left, key);
            var t = Floor(x.Right, key);
            return t ?? x;
        }

        public TKey Ceiling(TKey key) {
            var x = Ceiling(_root, key);
            return x == null ? default(TKey) : x.Key;
        }

        private TreeNode<TKey, TValue> Ceiling(TreeNode<TKey, TValue> x, TKey key) {
            if (x == null) return null;
            var cmp = key.CompareTo(x.Key);
            if (cmp == 1) return x;
            if (cmp > 0) return Ceiling(x.Right, key);
            var t = Ceiling(x.Left, key);
            return t ?? x;
        }

        public TKey Select(int k) {
            return Select(_root, k).Key;
        }

        private TreeNode<TKey, TValue> Select(TreeNode<TKey, TValue> x, int k) {
            if (x == null) return null;
            var t = Size(x.Left);
            if (t > k) return Select(x.Left, k);
            if (t < k) return Select(x.Right, k - t - 1);
            return x;
        }

        public int Rank(TKey key) {
            return Rank(key, _root);
        }

        private int Rank(TKey key, TreeNode<TKey, TValue> x) {
            if (x == null) return 0;
            var cmp = key.CompareTo(x.Key);
            if (cmp < 0) return Rank(key, x.Left);
            if (cmp > 0) return 1 + Size(x.Left) + Rank(key, x.Right);
            return Size(x.Left);
        }

        public IEnumerable<TKey> Keys() {
            return Keys(Min(), Max());
        }

        private IEnumerable<TKey> Keys(TKey lo, TKey hi) {
            var queue = new Queue<TKey>();
            Keys(_root, queue, lo, hi);
            return queue;
        }

        private void Keys(TreeNode<TKey, TValue> x, Queue<TKey> queue, TKey lo, TKey hi) {
            if (x == null) return;
            var cmplo = lo.CompareTo(x.Key);
            var cmphi = hi.CompareTo(x.Key);
            if (cmplo < 0) Keys(x.Left, queue, lo, hi);
            if (cmplo <= 0 && cmphi >= 0) queue.Enqueue(x.Key);
            if (cmphi > 0) Keys(x.Right, queue, lo, hi);
        }

        private bool IsEmpty() {
            return _root.N == 1;
        }

        private TreeNode<TKey, TValue> MoveRedLeft(TreeNode<TKey, TValue> h) {
            FlipColorsDesc(h);
            if (IsRed(h.Right.Left)) {
                h.Right = RotateRight(h.Right);
                h = RotateLeft(h);
            }
            return h;
        }

        public void DeleteMin() {
            if (!IsRed(_root.Left) && !IsRed(_root.Right))
                _root.color = RED;
            _root = DeleteMin(_root);
            if (!IsEmpty()) _root.color = BLACK;
        }

        private TreeNode<TKey, TValue> DeleteMin(TreeNode<TKey, TValue> x) {
            if (x.Left == null) return null;
            if (!IsRed(x.Left) && !IsRed(x.Left.Left))
                x = MoveRedLeft(x);
            x.Left = DeleteMin(x.Left);
            return Balance(x);
        }

        private TreeNode<TKey, TValue> MoveRedRight(TreeNode<TKey, TValue> h) {
            FlipColorsDesc(h);
            if (IsRed(h.Left.Left)) 
                h = RotateRight(h);
            return h;
        }

        public void DeleteMax() {
            if (!IsRed(_root.Left) && !IsRed(_root.Right))
                _root.color = RED;
            _root = DeleteMax(_root);
            if (!IsEmpty()) _root.color = BLACK;
        }

        private TreeNode<TKey, TValue> DeleteMax(TreeNode<TKey, TValue> h) {
            if (IsRed(h.Left))
                h = RotateRight(h);
            if (h.Right == null)
                return null;
            if (!IsRed(h.Right) && !IsRed(h.Right.Left))
                h = MoveRedRight(h);
            h.Right = DeleteMax(h.Right);
            return Balance(h);
        }

        public void Delete(TKey key) {
            if (!IsRed(_root.Left) && !IsRed(_root.Right))
                _root.color = RED;
            _root = Delete(_root, key);
            if (!IsEmpty()) _root.color = BLACK;
        }

        private TreeNode<TKey, TValue> Delete(TreeNode<TKey, TValue> h, TKey key) {
            var cmp = key.CompareTo(h.Key);
            if (cmp < 0) {
                if (!IsRed(h.Left) && !IsRed(h.Left.Left))
                    MoveRedLeft(h);
                h.Left = Delete(h.Left, key);
            }
            else {
                if (IsRed(h.Left))
                    h = RotateRight(h);
                if (cmp == 0 && h.Right == null)
                    return null;
                if (!IsRed(h.Right) && !IsRed(h.Right.Left))
                    h = MoveRedRight(h);
                if (cmp == 0) {
                    h.Value = Get(h.Right, Min(h.Right).Key);
                    h.Key = Min(h.Right).Key;
                    h.Right = DeleteMin(h.Right);
                }
                else {
                    h.Right = Delete(h.Right, key);
                }
            }
            return Balance(h);
        }
        
        private TreeNode<TKey, TValue> RotateLeft(TreeNode<TKey, TValue> h) {
            var x = h.Right;
            h.Right = x.Left;
            x.Left = h;
            x.color = h.color;
            h.color = RED;
            x.N = h.N;
            h.N = Size(h.Left) + 1 + Size(h.Right);
            return x;
        }

        private TreeNode<TKey, TValue> RotateRight(TreeNode<TKey, TValue> h) {
            var x = h.Left;
            h.Left = x.Right;
            x.Right = h;
            x.color = h.color;
            h.color = RED;
            x.N = h.N;
            h.N = Size(h.Left) + 1 + Size(h.Right);
            return x;
        }
        
        
    }
}