using System;

namespace LeetCodeTop100CSharp {
    class _5_LongestPalindromicSubstring {
        public string LongestPalindrome(string s) {
            if (s == null || s.Length < 1) {
                return "";
            }

            var start = 0;
            var end = 0;

            for (int i = 0; i < s.Length; i++) {
                var len1 = ExpandAroundCenter(s, i, i);
                var len2 = ExpandAroundCenter(s, i, i+1);
                var len = Math.Max(len1, len2);
                if(len > end - start) {
                    start = i - (len - 1) / 2;
                    end = i + len / 2;
                }
            }

            return s.Substring(start, end - start + 1);
        }

        private int ExpandAroundCenter(string s, int left, int right) {
            int L = left;
            int R = right;
            while (L >= 0 && R < s.Length && s[L] == s[R]) {
                L--;
                R++;
            }
            return R - L - 1;
        }
    }
}
