using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCodeTop100CSharp {
    class _26_RemoveDuplicatesfromSortedArray {
        public int RemoveDuplicates(int[] nums) {
            if(nums == null || nums.Length == 0) {
                return 0;
            }
            var slow = 0;

            for (int fast = 1; fast < nums.Length; fast++) {
                if(nums[fast] != nums[slow]) {
                    slow++;
                    nums[slow] = nums[fast];
                }
            }

            return slow + 1;
        }
    }
}
