using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCodeTop100CSharp {
    class _49_GroupAnagrams {

        public IList<IList<string>> GroupAnagrams(string[] strs) {
            var dic = new Dictionary<string, List<string>>();
            for (int i = 0; i < strs.Length; i++) {
                var origin = strs[i];
                var chars = strs[i].ToCharArray();
                Array.Sort(chars);
                var element = new string(chars);
                if (!dic.ContainsKey(element)) {
                    dic.Add(element, new List<string>() {origin});
                } else {
                    dic[element].Add(origin);
                }
            }

            var result = new string[dic.Count][];

            var index = 0;
            foreach (var item in dic) {
                result[index] = item.Value.ToArray();
                index++;
            }

            return result;
        }
    }
}
