using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCodeTop100CSharp.剑指Offer2 {
    class _0994_RottingOranges {
        public int OrangesRotting(int[][] grid) {
            if (grid == null || grid.Length == 0 || grid[0].Length == 0) {
                return -1;
            }

            var height = grid.Length ;
            var width = grid[0].Length ;

            var marked = new bool[height * width];
            var orangesCount = 0;
            var queue = new Queue<int[]>();

            for (int y = 0; y < height; y++) {
                for (int x = 0; x < width; x++) {
                    if (grid[y][x] == 2) {
                        marked[GetIndex(y, x, width)] = true;
                        queue.Enqueue(new int[] { y, x });
                    } else if (grid[y][x] == 0) {
                        marked[GetIndex(y, x, width)] = true;
                    } else {
                        orangesCount ++;
                    }
                }
            }

            if (queue.Count == 0 && orangesCount == 0)
                return 0;
            if (queue.Count == 0 && orangesCount > 0)
                return -1;

            var time = -1;
            
            while (queue.Count > 0) {
                var len = queue.Count;
                while (len > 0) {
                    var idx = queue.Dequeue();
                    //上
                    if (idx[0] - 1 >= 0 && !marked[GetIndex(idx[0] - 1, idx[1], width)]) {
                        marked[GetIndex(idx[0] - 1, idx[1], width)] = true;
                        queue.Enqueue(new int[] { idx[0] - 1, idx[1] });
                    }
                    //下
                    if (idx[0] + 1 < height && !marked[GetIndex(idx[0] + 1, idx[1], width)]) {
                        marked[GetIndex(idx[0] + 1, idx[1], width)] = true;
                        queue.Enqueue(new int[] { idx[0] + 1, idx[1] });
                    }
                    //左
                    if (idx[1] - 1 >= 0 && !marked[GetIndex(idx[0], idx[1] - 1, width)]) {
                        marked[GetIndex(idx[0], idx[1] - 1, width)] = true;
                        queue.Enqueue(new int[] { idx[0], idx[1] - 1});
                    }
                    //右
                    if (idx[1] + 1 < width && !marked[GetIndex(idx[0], idx[1] + 1, width)]) {
                        marked[GetIndex(idx[0], idx[1] + 1, width)] = true;
                        queue.Enqueue(new int[] { idx[0], idx[1] + 1});
                    }
                    len--;
                }
                time++;
            }

            foreach (var item in marked) {
                if (!item)
                    return -1;
            }

            return time;
        }

        public int GetIndex(int y, int x, int width) {
            return y * width + x;
        }


    }
}
