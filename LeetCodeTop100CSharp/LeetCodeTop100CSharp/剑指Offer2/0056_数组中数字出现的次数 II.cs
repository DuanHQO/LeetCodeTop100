using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCodeTop100CSharp.剑指Offer2 {
    class _0056_数组中数字出现的次数_II {
        public int SingleNumber(int[] nums) {

            var map = new Dictionary<int, int>();

            foreach (var item in nums) {
                if (map.ContainsKey(item))
                    map[item]++;
                else
                    map.Add(item, 1);
            }

            foreach (var item in map) {
                if (item.Value == 1)
                    return item.Key;
            }

            return 0;
        }
    }
}
