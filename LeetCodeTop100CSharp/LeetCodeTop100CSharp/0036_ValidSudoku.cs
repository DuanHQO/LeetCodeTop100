using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCodeTop100CSharp {
    class _0036_ValidSudoku {
        public bool IsValidSudoku(char[][] board) {
            var rowMap = new Dictionary<char, int>[9];
            var columnMap = new Dictionary<char, int>[9];
            var boxMap = new Dictionary<char, int>[9];

            for (int i = 0; i < 9; i++) {
                rowMap[i] = new Dictionary<char, int>();
                columnMap[i] = new Dictionary<char, int>();
                boxMap[i] = new Dictionary<char, int>();
            }

            for (int y = 0; y < 9; y++) {
                for (int x = 0; x < 9; x++) {
                    var alp = board[y][x];
                    var boxIdx = (y / 3) * 3 + x / 3;
                    if (alp != '.') {
                        if (rowMap[y].ContainsKey(alp)) {
                            Console.WriteLine("y {0} x {1} char {2}", y, x, alp);
                            return false;
                        } else {
                            rowMap[y].Add(alp, 1);
                        }
                        if (columnMap[x].ContainsKey(alp)) {
                            Console.WriteLine("y {0} x {1} char {2}", y, x, alp);
                            return false;
                        } else {
                            columnMap[x].Add(alp, 1);
                        }
                        if (boxMap[boxIdx].ContainsKey(alp)) {
                            Console.WriteLine("y {0} x {1} char {2}", y, x, alp);
                            return false;
                        } else {
                            boxMap[boxIdx].Add(alp, 1);
                        }
                    }
                }
            }

            return true;
        }
    }
}
