using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCodeTop100CSharp {
    class _461_HammingDistance {
        public static int HammingDistance(int x, int y) {
            var a = x ^ y;
            var dis = 0;
            while (a > 0) {
                if ((a & 1) == 1) {
                    dis++;
                }
                a >>= 1;
            }
            return dis;
        }
    }
}
