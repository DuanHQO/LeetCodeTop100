using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCodeTop100CSharp {
    class _35_SearchInsertPosition {
        public int SearchInsert(int[] nums, int target) {
            if(nums == null || nums.Length == 0) {
                return 0;
            }

            var lo = 0;
            var hi = nums.Length - 1;
            while (lo <= hi) {
                var mid = lo + (hi - lo) / 2;
                if (target == nums[mid]) {
                    return mid;
                } else if (target > nums[mid]) {
                    lo = mid + 1;
                } else {
                    hi = mid - 1;
                }
            }

            return lo;
            
        }
    }
}
