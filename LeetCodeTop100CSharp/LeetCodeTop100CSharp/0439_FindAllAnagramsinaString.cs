using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCodeTop100CSharp {
    class _439_FindAllAnagramsinaString {
        public static IList<int> FindAnagrams(string s, string p) {
            var length = s.Length;
            var res = new List<int>();

            var need = new Dictionary<char, int>();
            var window = new Dictionary<char, int>();

            foreach (var item in p) {
                if (need.ContainsKey(item)) {
                    need[item]++;
                } else {
                    need.Add(item, 1);
                }
            }

            var left = 0;
            var right = 0;
            var match = 0;

            while (right < length) {
                var chaR = s[right];
                if (need.ContainsKey(chaR)) {
                    if (window.ContainsKey(chaR)) {
                        window[chaR]++;
                    } else {
                        window.Add(chaR, 1);
                    }

                    if(need[chaR] == window[chaR]) {
                        match++;
                    }
                }
                right++;

                while (match == need.Count) {
                    if ((right - left) == p.Length) {
                        res.Add(left);
                    }
                    var chaL = s[left];
                    if (need.ContainsKey(chaL)) {
                        window[chaL]--;
                        if (window[chaL] < need[chaL]) {
                            match--;
                        }
                    }
                    left++;
                }
            }

            foreach (var item in res) {
                Console.WriteLine(item);
            }

            return res;
        }

        public static void Main(string[] args) {
            FindAnagrams("cbaebabacd", "abc");
            Console.Read();
        }
    }
}
