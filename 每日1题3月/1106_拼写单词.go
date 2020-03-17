package 每日1题3月

import "fmt"

func CountCharacters(words []string, chars string) int {
	if words == nil || len(words) == 0 {
		return 0
	}

	ret := 0

	need := make(map[byte]int)
	for i := 0; i < len(chars); i++ {
		char := chars[i]
		if _, ok := need[char]; ok {
			need[char]++
		} else {
			need[char] = 1
		}
	}

	for _, word := range words {
		tmp := make(map[byte]int)
		valid := true
		for i := 0; i < len(word); i++ {
			if _, ok := need[word[i]]; !ok {
				valid = false
				break
			}
			if _, ok := tmp[word[i]]; ok {
				tmp[word[i]]++
			} else {
				tmp[word[i]] = 1
			}
		}

		for key, value := range tmp {
			if value > need[key] {
				valid = false
				break
			}
		}

		if valid {
			ret += len(word)
		}
	}

	fmt.Println(ret)

	return ret
}
