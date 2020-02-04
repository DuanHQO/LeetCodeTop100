using System;
using System.Collections.Generic;

namespace LeetCodeTop100CSharp {
    class Program {
        static void Main(string[] args) {

            var test = new int?[] { 0, 2, 4, 1, null, 3, -1, 5, 1, null, 6, null, 8};
            var nodes = new TreeNode[test.Length];
            for (int i = 0; i < test.Length; i++) {
                if(test[i] != null) {
                    nodes[i] = new TreeNode(test[i].Value);
                }
            }

            var slow = 0;
            var fast = 1;

            while (fast < nodes.Length && fast + 1 < nodes.Length) {
                if (nodes[slow] != null) {
                    nodes[slow].left = nodes[fast];
                    nodes[slow].right = nodes[fast + 1];
                    //Console.WriteLine("node[{0}] left:{1} right:{2}", test[slow], test[fast], test[fast+1]);
                    fast += 2;
                }
                slow++;
            }

            
            foreach (var item in nodes) {
                if (item != null) {
                    //Console.WriteLine(item.val);
                }
            }
            Console.WriteLine(_104_MaximumDepthofBinaryTree.MaxDepth(nodes[0])); 

            Console.Read();
        }
    }

    public class TreeNode {
        public int val;
        public TreeNode left;
        public TreeNode right;
        public TreeNode(int x) { val = x; }
    }
}
