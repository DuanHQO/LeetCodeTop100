package 剑指Offer2

func exist(board [][]byte, word string) bool {
	if board == nil || len(board) == 0 || len(board[0]) == 0 {
		return false
	}

	var dfs func(x, y, k int) bool
	dfs = func(x, y, k int) bool {
		if y < 0 || y == len(board) || x < 0 || x == len(board[0]) || board[y][x] != word[k] {
			return false
		}

		if k == len(word)-1 {
			return true
		}

		tmp := board[y][x]
		board[y][x] = '/'
		res := dfs(x, y-1, k+1) || dfs(x, y+1, k+1) || dfs(x-1, y, k+1) || dfs(x+1, y, k+1)
		board[y][x] = tmp

		return res
	}

	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[0]); x++ {
			if dfs(x, y, 0) {
				return true
			}
		}
	}

	return false
}
