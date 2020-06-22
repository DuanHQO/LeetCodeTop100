using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCodeTop100CSharp {
    class SkipList<T> {
        private static readonly int DEFAULT_CAP = 65536;
        private int cap = 0;
        private int MaxLevel = 8;
        private readonly double p = 0.25;
        private SkipNode head;

        internal class SkipNode {
            public SkipNode[] forward;
            public int key;
            public T value;
        }

        public SkipList(int cap) {
            this.cap = cap;
            MaxLevel = SetMaxLevel(cap);

            head = new SkipNode();
            head.forward = new SkipNode[MaxLevel];
        }

        public SkipList() : this(DEFAULT_CAP) {
            
        }

        private int SetMaxLevel(int cap) {
            var level = Math.Log(cap) / Math.Log((1 / p));
            return (int)Math.Ceiling(level);
        }

        public int GetMaxLevel() {
            return this.MaxLevel;
        }

        public T Get(int key) {
            var cur = head;

            for (int l = MaxLevel - 1; l >= 0; l--) {
                if (cur.forward[l] == null) {
                    continue;
                }

                while (cur.forward[l] != null && key > cur.forward[l].key) {
                    cur = cur.forward[l];
                }

                if (cur.forward[l] != null && key == cur.forward[l].key) {
                    return cur.forward[l].value;
                }
            }

            return default(T);
        }

        public bool Set(int key, T value) {
            var node = InitNode(key, value);
            var newNodeMaxLevel = node.forward.Length;

            var cur = head;
            var update = new SkipNode[newNodeMaxLevel];

            for (int l = MaxLevel - 1; l >= 0; l--) {
                if (cur.forward[l] == null) {
                    if (l < newNodeMaxLevel) { 
                        update[l] = cur;
                    }
                    continue;
                }

                while (cur.forward[l] != null && key > cur.forward[l].key) {
                    cur = cur.forward[l];
                }

                if (cur.forward[l] != null && key == cur.forward[l].key) {
                    return false;
                }
                if (l < newNodeMaxLevel) {
                    update[l] = cur;
                }
            }

            for (int l = newNodeMaxLevel - 1; l >= 0; l--) {
                node.forward[l] = update[l].forward[l];
                update[l].forward[l] = node;
            }

            return true;
        }

        public bool Remove(int key) {
            var cur = head;
            var update = new SkipNode[MaxLevel];
            SkipNode delete = null;

            for (int l = MaxLevel - 1; l >=0; l--) {
                if (cur.forward[l] == null) {
                    continue;
                }

                while (cur.forward[l] != null && key > cur.forward[l].key) {
                    cur = cur.forward[l];
                }

                if (cur.forward[l] != null && key == cur.forward[l].key) {
                    delete = cur.forward[l];
                    update[l] = cur;
                }
            }

            if (delete == null) {
                return false;
            }

            for (int l = delete.forward.Length; l >= 0; l--) {
                update[l].forward[l] = delete.forward[l];
                delete.forward[l] = null;
            }


            return true;
        }

        public bool IsEmpty() {
            return head.forward[0] == null;
        }

        private int RandomLevel() {
            var level = 1;
            var rd = new Random();
            while (rd.NextDouble() <= p && level < MaxLevel) {
                level++;
            }
            return level;
        }

        private SkipNode InitNode(int key, T value) {
            var node = new SkipNode {
                key = key,
                value = value
            };
            var level = RandomLevel();
            node.forward = new SkipNode[level];
            return node;
        }


        public override string ToString() {
            SkipNode cur = null;
            var sb = new StringBuilder();
            for (int l = MaxLevel; l >= 0; l--) {
                sb.Append("第").Append(l).Append("层");
                cur = head.forward[l];

                while (cur != null) {
                    sb.Append(cur.key).Append(";");
                    cur = cur.forward[l];
                }
                sb.Append("\n");
            }
            return sb.ToString();
        }
    }
}
