using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCodeTop100CSharp {
    class _126_SingleNumber {
        public int SingleNumber(int[] nums) {
            var tmp = nums[0];
            for (int i = 1; i < nums.Length; i++) {
                tmp ^= nums[i];
            }
            return tmp;
        }
    }
}
