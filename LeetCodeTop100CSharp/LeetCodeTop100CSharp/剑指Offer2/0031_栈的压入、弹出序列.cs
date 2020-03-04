using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCodeTop100CSharp.剑指Offer2 {
    class _0031_栈的压入_弹出序列 {
        public bool ValidateStackSequences(int[] pushed, int[] popped) {
            var len = pushed.Length;
            var j = 0;
            var stack = new Stack<int>();
            for (int i = 0; i < len; i++) {
                stack.Push(pushed[i]);
                while (stack.Count > 0 && stack.Peek() == popped[j]) {
                    j++;
                    stack.Pop();
                }
            }

            return len == j;
        }
    }
}
