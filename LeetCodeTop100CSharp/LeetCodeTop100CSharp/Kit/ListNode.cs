namespace LeetCodeTop100CSharp {
    
    public class ListNode{
        public int Val;
        public ListNode Left;
        public ListNode Right;
        public ListNode Next;
        public ListNode Random;

        public ListNode() { }

        public ListNode(int val) {
            Val = val;
        }

        public ListNode(int val, ListNode left, ListNode right, ListNode next) {
            Val = val;
            Left = left;
            Right = right;
            Next = next;
        }
    }
}