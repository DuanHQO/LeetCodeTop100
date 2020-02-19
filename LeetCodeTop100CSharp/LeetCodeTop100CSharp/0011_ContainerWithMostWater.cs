using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCodeTop100CSharp {
    class _11_ContainerWithMostWater {

        public static int MaxArea(int[] height) {
            var i = 0;
            var j = height.Length - 1;
            var res = 0;

            while (i < j) {
                if (height[i] < height[j]) {
                    res = Math.Max(res, (j - i) * height[i++]);
                } else {
                    res = Math.Max(res, (j - i) * height[j--]);
                }
            }

            return res;
        }
    }
}
