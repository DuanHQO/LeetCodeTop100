using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCodeTop100CSharp.剑指Offer2 {
    class _0035_复杂链表的复制 {
        public Node CopyRandomList(Node head) {
            if (head == null)
                return null;

            var map = new Dictionary<Node, Node>();

            var node = head;
            while (node != null) {
                if (!map.ContainsKey(node)) {
                    var tmp = new Node();
                    tmp.val = node.val;
                    map.Add(node, tmp);
                }

                if (node.next != null) {
                    if (!map.ContainsKey(node.next)) {
                        var tmp = new Node();
                        tmp.val = node.next.val;
                        map.Add(node.next, tmp);
                    }
                    map[node].next = map[node.next];
                }

                if (node.random != null) {
                    if (!map.ContainsKey(node.random)) {
                        var tmp = new Node();
                        tmp.val = node.random.val;
                        map.Add(node.random, tmp);
                    }
                    map[node].random = map[node.random];
                }
                node = node.next;
            }

            return map[head];
        }
    }
}
