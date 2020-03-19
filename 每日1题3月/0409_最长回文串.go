package 每日1题3月

import "fmt"

func LongestPalindrome(s string) int {
	if s == "" {
		return 0
	}

	ret := 0

	dic := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		char := s[i]
		if _, ok := dic[char]; ok {
			dic[char]++
		} else {
			dic[char] = 1
		}
	}

	for _, value := range dic {
		if (ret&1 == 1) && (value&1 == 1) {
			ret += value - 1
		} else {
			ret += value
		}
	}

	fmt.Println(ret)
	return ret

}
