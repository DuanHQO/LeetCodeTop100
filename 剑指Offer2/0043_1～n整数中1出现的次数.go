package 剑指Offer2

func countDigitOne(n int) int {
	pre := 0
	for i := 1; i <= n; i++ {
		count := 0
		tmp := i
		for tmp > 0 {
			l := tmp % 10
			if l == 1 {
				count++
			}
			tmp /= 10
		}

		pre += count
	}

	return pre
}
