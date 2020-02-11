using System;
using System.Collections.Generic;
using System.Text;

namespace LeetCodeTop100CSharp {
    class _394_DecodeString {
        public static string DecodeString(string s) {
            var str = new StringBuilder("");

            var numStack = new Stack<int>();
            var alpStack = new Stack<string>();

            var multi = 0;

            foreach (var item in s.ToCharArray()) {
                if (item == '[') {
                    numStack.Push(multi);
                    alpStack.Push(str.ToString());
                    multi = 0;
                    str.Clear();
                } else if (item == ']') {
                    var tmp = new StringBuilder();
                    var curMulti = numStack.Pop();
                    for (int i = 0; i < curMulti; i++) {
                        tmp.Append(str);
                    }
                    str = new StringBuilder(alpStack.Pop() + tmp);
                } else if (item >= '0' && item <= '9') {
                    multi = multi * 10 + int.Parse(item.ToString());
                } else {
                    str.Append(item);
                }
            }

            return str.ToString();
        }

        public static void Main(string[] args) {
            Console.WriteLine(DecodeString("3[a]2[bc]"));
            Console.WriteLine(DecodeString("3[a2[c]]"));
            Console.WriteLine(DecodeString("2[abc]3[cd]ef"));
            Console.Read();
        }
    }
}
