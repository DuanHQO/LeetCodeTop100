using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCodeTop100CSharp {
    class _142_LinkedListCycleII {

        public ListNode DetectCycle(ListNode head) {
            if (head == null || head.Next == null) {
                return null;
            }
            var dic = new Dictionary<ListNode, ListNode>();

            var pre = head;

            dic.Add(head, head);
            head = head.Next;

            while (head != null) {
                if(dic.ContainsKey(head)) {
                    return head;
                } else {
                    dic.Add(head, pre);
                    pre = pre.Next;
                    head = head.Next;
                }
            }

            return null;
        }
    }
}
