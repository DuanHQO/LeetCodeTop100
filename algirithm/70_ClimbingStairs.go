package algirithm

func ClimbStairs(n int) int {
	if n <= 2 {
		return n
	}

	i1 := 1
	i2 := 2
	for i := 3; i <= n; i++ {
		temp := i1 + i2
		i1 = i2
		i2 = temp
	}
	return i2
}
