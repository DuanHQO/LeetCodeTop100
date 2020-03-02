package 剑指Offer2

func constructArr(a []int) []int {
	if a == nil {
		return nil
	}

	n := len(a)
	L, R := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		L[i] = 1
		R[i] = 1
	}

	for i := 1; i < n; i++ {
		L[i] = L[i-1] * a[i-1]
	}

	for i := n - 2; i >= 0; i-- {
		R[i] = R[i+1] * a[i+1]
	}

	for i := 0; i < n; i++ {
		L[i] = L[i] * R[i]
	}

	return L
}
