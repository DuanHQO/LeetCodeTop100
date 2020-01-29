using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCodeTop100CSharp {
    class _169_MajorityElement {
        public int MajorityElement(int[] nums) {
            var count = 0;
            var major = 0;
            for (int i = 0; i < nums.Length; i++) {
                if(count == 0) {
                    major = nums[i];
                    count = 1;
                } else {
                    count += (nums[i] == major) ? 1 : -1;
                }
            }
            return major;
        }
    }
}
