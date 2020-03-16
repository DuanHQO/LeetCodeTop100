using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCodeTop100CSharp {
    class _27_RemoveElement {
        public int RemoveElement(int[] nums, int val) {
            if(nums == null || nums.Length == 0) {
                return 0;
            }
            var slow = 0;
            for (int fast = 0; fast < nums.Length; fast++) {
                if(nums[fast] != val) {
                    nums[slow] = nums[fast];
                    slow++;
                }
            }
            return slow;
        }
    }
}
