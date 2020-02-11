using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCodeTop100CSharp {
    class _139_WordBreak {

        public bool WordBreak(string s, IList<string> wordDict) {
            var wordDictSet = new HashSet<string>(wordDict);
            var queue = new Queue<int>();
            var visited = new int[s.Length];
            queue.Enqueue(0);

            while (queue.Count > 0) {
                var start = queue.Dequeue();
                if (visited[start] == 0) {
                    for (int end = start + 1; end <= s.Length; end++) {
                        if(wordDictSet.Contains(s.Substring(start, end - start))) {
                            queue.Enqueue(end);
                            if(end == s.Length) {
                                return true;
                            }
                        }
                    }
                    visited[start] = 1;
                }
            }

            return false;
        }
    }
}
