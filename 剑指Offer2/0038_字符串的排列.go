package 剑指Offer2

import (
	"fmt"
	"sort"
)

func Permutation(s string) []string {
	var res []string
	if s == "" {
		return res
	}

	char := chars(s)

	sort.Sort(char)

	var helper func(s []byte, tmp []byte)
	helper = func(s []byte, tmp []byte) {
		if len(tmp) == len(char) {
			slice := make([]byte, len(tmp), len(tmp))
			copy(slice, tmp)
			res = append(res, string(slice))
			return
		}

		for i := 0; i < len(s); i++ {
			if i > 0 && s[i] == s[i-1] {
				continue
			}
			//marked[i] = true
			tmp = append(tmp, s[i])
			ss := make([]byte, len(s), len(s))
			copy(ss, s)
			helper(append(ss[:i], ss[i+1:]...), tmp)
			tmp = tmp[:len(tmp)-1]
		}

	}

	helper(char, []byte{})
	fmt.Println(res)
	return res
}
