package algirithm

func CountBinarySubstrings(s string) int {
	ret := 0
	if len(s) == 0 {
		return ret
	}

	var min func(int, int) int
	min = func(a int, b int) int {
		if a < b {
			return a
		}
		return b
	}

	last := 0
	ptr := 0
	for ptr < len(s) {
		char := s[ptr]
		count := 0
		for ptr < len(s) && char == s[ptr] {
			count++
			ptr++
		}
		ret += min(last, count)
		last = count
	}

	return ret
}
