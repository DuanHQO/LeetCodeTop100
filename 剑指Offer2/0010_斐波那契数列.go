package 剑指Offer2

func fib(n int) int {

	a := 0
	b := 1

	for i := 1; i <= n; i++ {
		a, b = b%1000000007, a+b
	}

	return a
}
