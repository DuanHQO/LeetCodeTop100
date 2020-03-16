using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCodeTop100CSharp {
    class _406_QueueReconstructionbyHeight {

        public static int[][] ReconstructQueue(int[][] people) {

            var result = new List<int[]>();

            Array.Sort(people, (a, b) => {
                return a[0] == b[0] ? a[1] - b[1] : b[0] - a[0];
            });

            for (int i = 0; i < people.Length; i++) {
                Console.WriteLine("[{0}, {1}]", people[i][0], people[i][1]);
                result.Insert(people[i][1], people[i]);
            }
            Console.WriteLine();
            foreach (var item in result) {
                Console.WriteLine("[{0}, {1}]", item[0], item[1]);
            }

            return result.ToArray();
        }

        public static void Main(string[] args) {
            ReconstructQueue(new int[][] {
                new int[]{7, 0 },  new int[]{4, 4 },  new int[] { 7, 1 },  new int[]{5, 0 },  new int[]{6, 1 },  new int[]{5, 2 }
            });
            Console.Read();
        }
    }
}
