using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCodeTop100CSharp.剑指Offer2 {
    class _0060_n个骰子的点数 {
        public static double[] TwoSum(int n) {
            var dp = new int[n + 1][];
            for (int i = 1; i < dp.Length; i++) {
                dp[i] = new int[11 * 6 + 1];
            }

            for (int i = 1; i <= 6; i++) {
                dp[1][i] = 1;
            }

            for (int i = 2; i <= n; i++) {
                for (int j = i; j <= 6*i; j++) {
                    for (int cur = 1; cur <= 6; cur++) {
                        if (j-cur <= 0) {
                            break;
                        }
                        dp[i][j] += dp[i - 1][j - cur];
                    }
                }
            }

            var total = Math.Pow(6, n);
            var ret = new List<double>();
            for (int i = n; i <= 6 *n; i++) {
                ret.Add(dp[n][i] / total);
            }

            return ret.ToArray();
        }

        public static void Main(string[] args) {
            TwoSum(2);
        }
    }
}
