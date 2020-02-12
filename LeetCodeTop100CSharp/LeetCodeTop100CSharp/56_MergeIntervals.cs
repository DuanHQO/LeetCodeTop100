using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCodeTop100CSharp {
    class _56_MergeIntervals {
        public static int[][] Merge(int[][] intervals) {
            if (intervals == null ) {
                return null;
            }

            if(intervals.Length <= 1) {
                return intervals;
            }

            Array.Sort(intervals, (a, b) => {
                return a[0] == b[0] ? -1 : a[0] - b[0];
            });

            foreach (var item in intervals) {
                Console.WriteLine(item[0] + " " + item[1]);
            }

            Console.WriteLine();

            var res = new List<int[]>();
            for (int i = 0; i < intervals.Length; i++) {
                Console.WriteLine(intervals[i][0]);
                if(res.Count == 0 || res.Last()[1] < intervals[i][0]) {
                    res.Add(intervals[i]);
                } else {
                    res.Last()[1] = Math.Max(res.Last()[1], intervals[i][1]);
                }
            }

            return res.ToArray();
        }

        public static void Main(string[] args) {
            Merge(new int[][] {
                new int[] { 2, 6 },
                new int[] { 1, 3 },
                new int[] { 15, 18 },
                new int[] { 8, 10 },
            });
        }
    }
}
