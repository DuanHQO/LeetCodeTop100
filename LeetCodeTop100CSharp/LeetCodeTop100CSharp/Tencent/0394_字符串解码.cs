using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCodeTop100CSharp.Tencent {
    class _0394_字符串解码 {
        public string DecodeString(string s) {
            var str = new StringBuilder("");

            var numStack = new Stack<int>();
            var alpStack = new Stack<string>();

            var mul = 0;


            foreach (var item in s) {
                if (item == '[') {
                    numStack.Push(mul);
                    alpStack.Push(str.ToString());
                    mul = 0;
                    str.Clear();
                } else if (item == ']') {
                    var tmp = new StringBuilder();
                    var curMul = numStack.Pop();
                    for (int i = 0; i < curMul; i++) {
                        tmp.Append(str);
                    }
                    str = new StringBuilder(alpStack.Pop() + tmp);
                } else if (item >= '0' && item <= '9') {
                    mul = mul * 10 + int.Parse(item.ToString());
                } else {
                    str.Append(item);
                }
            }

            return str.ToString();
        }
    }
}
