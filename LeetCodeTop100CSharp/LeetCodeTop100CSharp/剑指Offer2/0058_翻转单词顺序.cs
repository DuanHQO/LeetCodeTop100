using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCodeTop100CSharp.剑指Offer2 {
    class _0058_翻转单词顺序 {
        public static string ReverseWords(string s) {
            if (string.IsNullOrEmpty(s))
                return string.Empty;

            var arr = s.Split(new string[] {" "}, StringSplitOptions.RemoveEmptyEntries);
            var revarr = arr.Reverse();
            return string.Join(" ", revarr);
        }

        public static void Main(string[] args) {
            ReverseWords("the sky is blue");
            Console.Read();
        }
    }
}
