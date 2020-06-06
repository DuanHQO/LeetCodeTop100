using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCodeTop100CSharp {
    class _0138_CopyListwithRandomPointer {
        public ListNode CopyRandomList(ListNode head) {
            if (head == null) return null;
            var map = new Dictionary<ListNode, ListNode>();
            var node = head;
            while (node != null) {
                if (!map.ContainsKey(node)) {
                    map.Add(node, new ListNode(node.val));
                }
                if (node.next != null) { 
                    if (!map.ContainsKey(node.next)) 
                        map.Add(node.next, new ListNode(node.next.val));
                    map[node].next = map[node.next];
                }
                if (node.random != null ) {
                    if (!map.ContainsKey(node.random)) 
                        map.Add(node.random, new ListNode(node.random.val));
                    map[node].random = map[node.random];
                }
                node = node.next;
            }

            return map[head];
        }
    }
}
