using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCodeTop100CSharp {
    class _18_4Sum {
        public static IList<IList<int>> FourSum(int[] nums, int target) {
            if (nums == null) {
                return null;
            }

            var res = new List<int[]>();

            if (nums.Length == 0) {
                return res.ToArray();
            }

            Array.Sort(nums, (a, b) => {
                return a - b;
            });

            foreach (var item in nums) {
                Console.Write(item + " ");
            }

            //if (nums[0] > target) {
            //    return res.ToArray();
            //}

            for (int a = 0; a <= nums.Length - 4; a++) {
                if(a > 0 && nums[a] == nums[a - 1]) {
                    continue;
                }
                for (int b = a + 1; b <= nums.Length - 3; b++) {
                    if( b > a + 1 && nums[b] == nums[b - 1]) {
                        continue;
                    }
                    var c = b + 1;
                    var d = nums.Length - 1;

                    while (c < d) {
                        var sum = nums[a] + nums[b] + nums[c] + nums[d];

                        if(sum > target) {
                            d--;
                        } else if (sum < target) {
                            c++;
                        } else {
                            var tmp = new int[4] { nums[a], nums[b], nums[c], nums[d] };
                            res.Add(tmp);
                            c++;
                            d--;
                            while (c < d && nums[c] == nums[c-1]) {
                                c++;
                            }
                            while (c < d && nums[d] == nums[d+1]) {
                                d--;
                            }
                        }
                    }
                }
            }
            return res.ToArray();
        }

        public static void Main(string[] args) {
            var res = FourSum(new int[]{ 1,-2,-5,-4,-3,3,3,5}, -11);
        }
    }
}
