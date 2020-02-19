using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCodeTop100CSharp {
    class _647_PalindromicSubstrings {

        int num = 0;

        public int CountSubstrings(string s) {
            for (int i = 0; i < s.Length; i++) {
                count(s, i, i);
                count(s, i, i + 1);
            }
            return num;
        }

        public void count(string s, int start, int end) {
            while (start >= 0 && end < s.Length && s[start] == s[end]){
                num++;
                start--;
                end++;
            }
        }
    }
}
