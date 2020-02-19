package algirithm

func NumTrees(n int) int {
	G := make([]int, n+1)
	G[0] = 1
	G[1] = 1

	for count := 2; count <= n; count++ {
		for i := 1; i <= count; i++ {
			G[count] += G[i-1] * G[count-i]
		}
	}

	return G[n]
}
