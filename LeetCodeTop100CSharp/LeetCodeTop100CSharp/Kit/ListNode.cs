namespace LeetCodeTop100CSharp {
    
    public class ListNode{
        public int val;
        public ListNode left;
        public ListNode right;
        public ListNode next;
        public ListNode random;

        public ListNode() { }

        public ListNode(int _val) {
            val = _val;
        }

        public ListNode(int _val, ListNode _left, ListNode _right, ListNode _next) {
            val = _val;
            left = _left;
            right = _right;
            next = _next;
        }
    }
}