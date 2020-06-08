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
                    map.Add(node, new ListNode(node.Val));
                }
                if (node.Next != null) { 
                    if (!map.ContainsKey(node.Next)) 
                        map.Add(node.Next, new ListNode(node.Next.Val));
                    map[node].Next = map[node.Next];
                }
                if (node.Random != null ) {
                    if (!map.ContainsKey(node.Random)) 
                        map.Add(node.Random, new ListNode(node.Random.Val));
                    map[node].Random = map[node.Random];
                }
                node = node.Next;
            }

            return map[head];
        }
    }
}
