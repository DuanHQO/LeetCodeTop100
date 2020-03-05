package 剑指Offer2

func sumNums(n int) int {
	ret := 0

	var helper func(n int) bool
	helper = func(n int) bool {
		ret += n
		return n != 0 && helper(n-1)
	}

	helper(n)
	return ret
}
