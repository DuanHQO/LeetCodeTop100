using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCodeTop100CSharp {
    class _0138_CopyListwithRandomPointer {
        public Node CopyRandomList(Node head) {
            if (head == null) return null;
            var map = new Dictionary<Node, Node>();
            var node = head;
            while (node != null) {
                if (!map.ContainsKey(node)) {
                    map.Add(node, new Node(node.val));
                }
                if (node.next != null) { 
                    if (!map.ContainsKey(node.next)) 
                        map.Add(node.next, new Node(node.next.val));
                    map[node].next = map[node.next];
                }
                if (node.random != null ) {
                    if (!map.ContainsKey(node.random)) 
                        map.Add(node.random, new Node(node.random.val));
                    map[node].random = map[node.random];
                }
                node = node.next;
            }

            return map[head];
        }
    }
}
