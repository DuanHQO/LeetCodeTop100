using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCodeTop100CSharp.剑指Offer2 {
    class _0045_把数组排成最小的数 {
        public static string MinNumber(int[] nums) {
            if (nums == null || nums.Length == 0) {
                return string.Empty;
            }

            var numstr = new string[nums.Length];
            for (int i = 0; i < nums.Length; i++) {
                numstr[i] = nums[i].ToString();
            }

            for (int i = 0; i < nums.Length; i++) {
                for (int j = i + 1; j < nums.Length; j++) {
                    var sum = numstr[i] + numstr[j];
                    if (string.Compare(sum, numstr[j] + numstr[i]) > 0){
                        var tmps = numstr[i];
                        numstr[i] = numstr[j];
                        numstr[j] = tmps;
                    }
                }
            }
            return string.Join("", numstr);
        }

        public static void Main(string[] args) {
            MinNumber(new int[] { 3, 30, 34, 5, 9 });
            Console.Read();
        }
    }
}
