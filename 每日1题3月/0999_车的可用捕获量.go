package 每日1题3月

import "fmt"

func NumRookCaptures(board [][]byte) int {
	ret := 0

	var dfs func(y, x, dir int)
	dfs = func(y, x, dir int) {
		if y < 0 || y == len(board) || x < 0 || x == len(board[0]) || board[y][x] == 'B' {
			return
		}

		fmt.Printf("[%d, %d] = %c, dir = %d\n", y, x, board[y][x], dir)

		if board[y][x] == 'p' {
			ret++
			return
		}

		switch dir {
		case 0:
			dfs(y-1, x, dir)
		case 1:
			dfs(y+1, x, dir)
		case 2:
			dfs(y, x-1, dir)
		case 3:
			dfs(y, x+1, dir)
		}
	}

	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[0]); x++ {
			if board[y][x] == 'R' {
				dfs(y, x, 0)
				dfs(y, x, 1)
				dfs(y, x, 2)
				dfs(y, x, 3)
			}
		}
	}

	fmt.Println(ret)

	return ret
}
