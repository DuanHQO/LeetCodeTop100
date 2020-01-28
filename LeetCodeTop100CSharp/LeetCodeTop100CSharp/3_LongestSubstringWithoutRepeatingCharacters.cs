using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCodeTop100CSharp {
    class _3_LongestSubstringWithoutRepeatingCharacters {
        public static int LengthOfLongestSubstring(string s) {

            if (string.IsNullOrEmpty(s)) {
                return 0;
            }
            var chars = s.ToCharArray();
            var dic = new Dictionary<char, int>();

            var num = 0;
            var startindex = 0;
            var endindex = 0;


            for (int i = 0; i < chars.Length; i++) {
                if(dic.ContainsKey(chars[i])) {
                    startindex = Math.Max(dic[chars[i]], startindex);
                    dic[chars[i]] = i + 1;
                } else {
                    dic.Add(chars[i], i + 1);
                }
                endindex = i;
                num = Math.Max(endindex - startindex + 1, num);
            }
            return num;
        }

    }
}
