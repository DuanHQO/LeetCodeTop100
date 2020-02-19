package algirithm

func MinWindow(s string, t string) string {
	if s == "" || t == "" {
		return ""
	}

	need := make(map[byte]int)
	window := make(map[byte]int)

	for i := 0; i < len(t); i++ {
		_, ok := need[t[i]]
		if ok {
			need[t[i]]++
		} else {
			need[t[i]] = 1
		}
	}

	left := 0
	right := 0
	match := 0
	start := 1 << 31
	minLen := 1 << 31
	end := len(s)

	for ; right < len(s); right++ {
		char := s[right]
		_, ok := need[char]
		if ok {
			_, exists := window[char]
			if exists {
				window[char]++
			} else {
				window[char] = 1
			}
			if window[char] == need[char] {
				match++
			}
		}

		for ; match == len(need); left++ {
			if (right - left + 1) < minLen {
				minLen = right - left + 1
				start = left
				end = right + 1
			}

			char2 := s[left]
			_, exists := need[char2]
			if exists {
				window[char2]--
				if window[char2] < need[char2] {
					match--
				}
			}
		}
	}

	if minLen == 1<<31 {
		return ""
	} else {
		return s[start:end]
	}
}
