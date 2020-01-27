using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCodeTop100CSharp {
    class ValidParentheses {
        public static bool IsValid(string s) {
            var chars = s.ToCharArray();
            if ((chars.Length & 1) == 1) {
                return false;
            }
            var stack = new Stack<char>();
            foreach (var item in chars) {
                if (stack.Count == 0) {
                    stack.Push(item);
                    continue;
                }
                switch (item) {
                    case '(':
                        stack.Push(item);
                        break;
                    case ')':
                        if (stack.Peek() == '(') {
                            stack.Pop();
                        } else {
                            stack.Push(item);
                        }
                        break;
                    case '[':
                        stack.Push(item);
                        break;
                    case ']':
                        if (stack.Peek() == '[') {
                            stack.Pop();
                        } else {
                            stack.Push(item);
                        }
                        break;
                    case '{':
                        stack.Push(item);
                        break;
                    case '}':
                        if (stack.Peek() == '{') {
                            stack.Pop();
                        } else {
                            stack.Push(item);
                        }
                        break;
                }
            }
            if (stack.Count > 0) {
                return false;
            }
            return true;
        }
    }
}
