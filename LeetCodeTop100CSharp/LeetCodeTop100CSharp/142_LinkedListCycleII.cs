using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCodeTop100CSharp {
    class _142_LinkedListCycleII {

        public ListNode DetectCycle(ListNode head) {
            if (head == null || head.next == null) {
                return null;
            }
            var dic = new Dictionary<ListNode, ListNode>();

            var pre = head;

            dic.Add(head, head);
            head = head.next;

            while (head != null) {
                if(dic.ContainsKey(head)) {
                    return head;
                } else {
                    dic.Add(head, pre);
                    pre = pre.next;
                    head = head.next;
                }
            }

            return null;
        }
    }
}
