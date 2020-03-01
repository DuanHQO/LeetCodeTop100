package 剑指Offer2

func printNumbers(n int) []int {
	max := 1
	for i := 1; i <= n; i++ {
		max *= 10
	}

	var res []int
	for i := 1; i < max; i++ {
		res = append(res, i)
	}

	return res
}
