using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCodeTop100CSharp.剑指Offer2 {
    class _0035_复杂链表的复制 {
        public ListNode CopyRandomList(ListNode head) {
            if (head == null)
                return null;

            var map = new Dictionary<ListNode, ListNode>();

            var node = head;
            while (node != null) {
                if (!map.ContainsKey(node)) {
                    var tmp = new ListNode();
                    tmp.Val = node.Val;
                    map.Add(node, tmp);
                }

                if (node.Next != null) {
                    if (!map.ContainsKey(node.Next)) {
                        var tmp = new ListNode();
                        tmp.Val = node.Next.Val;
                        map.Add(node.Next, tmp);
                    }
                    map[node].Next = map[node.Next];
                }

                if (node.Random != null) {
                    if (!map.ContainsKey(node.Random)) {
                        var tmp = new ListNode();
                        tmp.Val = node.Random.Val;
                        map.Add(node.Random, tmp);
                    }
                    map[node].Random = map[node.Random];
                }
                node = node.Next;
            }

            return map[head];
        }
    }
}
