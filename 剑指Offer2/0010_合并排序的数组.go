package 剑指Offer2

func Merge0010(A []int, m int, B []int, n int) {
	var ret []int
	a, b := 0, 0
	for a < m && b < n {
		if A[a] < B[b] {
			ret = append(ret, A[a])
			a++
		} else {
			ret = append(ret, B[b])
			b++
		}
	}

	if a == m {
		ret = append(ret, B[b:]...)
	} else {
		ret = append(ret, A[a:m]...)
	}
	copy(A, ret)
}
