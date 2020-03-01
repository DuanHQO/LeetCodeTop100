package 剑指Offer2

func firstUniqChar(s string) byte {
	if s == "" {
		return ' '
	}

	nums := make([]int, 26)
	for _, value := range s {
		idx := value - 'a'
		nums[idx]++
	}

	for _, char := range s {
		if nums[char-'a'] == 1 {
			return byte(char)
		}
	}

	return ' '
}
