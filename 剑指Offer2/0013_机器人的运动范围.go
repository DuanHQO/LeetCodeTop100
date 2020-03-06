package 剑指Offer2

func movingCount(m int, n int, k int) int {

	var sums func(a int) int
	sums = func(a int) int {
		tmp := 0
		for a > 0 {
			tmp += a % 10
			a /= 10
		}
		return tmp
	}

	ret := 0

	marked := make([][]bool, m)
	for i := 0; i < len(marked); i++ {
		marked[i] = make([]bool, n)
	}

	var dfs func(x, y int)
	dfs = func(x, y int) {
		if y == m || x == n || marked[y][x] || sums(x)+sums(y) > k {
			if y < m && x < n {
				marked[y][x] = true
			}
			return
		}

		marked[y][x] = true
		ret++
		dfs(x, y+1)
		dfs(x+1, y)
	}

	dfs(0, 0)

	return ret
}
