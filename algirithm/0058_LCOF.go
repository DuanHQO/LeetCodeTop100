package algirithm

func reverseLeftWords(s string, n int) string {
	if s == "" || n < 0 || n > len(s) {
		return ""
	}

	left := s[:n]
	right := s[n:]

	return right + left
}
