package 剑指Offer2

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}

	ret := 0

	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	dic := make(map[byte]int)
	start := 0
	for i := 0; i < len(s); i++ {
		if _, ok := dic[s[i]]; ok {
			start = max(dic[s[i]], start)
		}
		ret = max(ret, i-start+1)
		dic[s[i]] = i + 1
	}

	return ret
}
