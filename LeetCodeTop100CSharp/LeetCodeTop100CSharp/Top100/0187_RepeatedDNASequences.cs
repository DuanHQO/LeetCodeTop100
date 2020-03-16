using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCodeTop100CSharp {
    class _0187_RepeatedDNASequences {
        public IList<string> FindRepeatedDnaSequences(string s) {
            var res = new HashSet<string>();
            if (s == null || s.Length < 10) {
                return res.ToArray();
            }

            var map = new Dictionary<string, int>();
            for (int i = 0; i < s.Length - 9; i++) {
                var sub = s.Substring(i, 10);
                if (!map.ContainsKey(sub)) {
                    map.Add(sub, 1);
                } else {
                    res.Add(sub);
                }
            }

            return res.ToArray();
        }
    }
}
