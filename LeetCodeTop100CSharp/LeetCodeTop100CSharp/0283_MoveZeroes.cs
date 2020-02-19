using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCodeTop100CSharp {
    class _283_MoveZeroes {
        public void MoveZeroes(int[] nums) {
            if (nums == null) {
                return;
            }
            if(nums.Length <= 1) {
                return;
            }
            var lt = 0;
            var i = 1;
            while (i < nums.Length) {
                if(nums[lt]!= 0) {
                    lt++;
                    i = lt + 1;
                } else {
                    if(nums[i] != 0) {
                        nums[lt] = nums[i];
                        nums[i] = 0;
                    }
                    i++;
                }
            }
        }
    }
}
