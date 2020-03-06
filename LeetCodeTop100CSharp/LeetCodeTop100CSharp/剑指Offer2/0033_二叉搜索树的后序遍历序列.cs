using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCodeTop100CSharp.剑指Offer2 {
    class _0033_二叉搜索树的后序遍历序列 {
        public bool VerifyPostorder(int[] postorder) {
            if (postorder == null || postorder.Length == 0)
                return true;

            var stack = new Stack<int>();
            var pre = int.MaxValue;
            for (int i = postorder.Length - 1; i >= 0 ; i--) {
                if (postorder[i] > pre)
                    return false;
                while (stack.Count > 0 && postorder[i] < stack.Peek()) {
                    pre = stack.Pop();
                }
                stack.Push(postorder[i]);
            }


            return true;
        }
    }
}
