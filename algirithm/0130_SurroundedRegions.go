package algirithm

import "fmt"

func Solve(board [][]byte) {
	fmt.Printf("%v\n", board)
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}

	height, width := len(board), len(board[0])

	var dfs func(y, x int)
	dfs = func(y, x int) {
		if y < 0 || y >= height || x < 0 || x >= width || board[y][x] != 'O' {
			return
		}
		fmt.Println(y, x)
		board[y][x] = 'A'
		dfs(y-1, x)
		dfs(y+1, x)
		dfs(y, x-1)
		dfs(y, x+1)
	}

	for i := 0; i < height; i++ {
		dfs(i, 0)
		dfs(i, width-1)
	}

	for j := 0; j < width; j++ {
		dfs(0, j)
		dfs(height-1, j)
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if board[y][x] == 'A' {
				board[y][x] = 'O'
			} else if board[y][x] == 'O' {
				board[y][x] = 'X'
			}
		}
	}
}
