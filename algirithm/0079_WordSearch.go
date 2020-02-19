package algirithm

import "fmt"

var markedExist [][]bool

func resetMarkedExist() {
	for y := 0; y < len(markedExist); y++ {
		for x := 0; x < len(markedExist[0]); x++ {
			markedExist[y][x] = false
		}
	}
}

func Exist(board [][]byte, word string) bool {

	if board == nil || word == "" {
		return false
	}

	dic := make(map[byte][][]int)
	markedExist = make([][]bool, len(board))
	for y := 0; y < len(board); y++ {
		markedExist[y] = make([]bool, len(board[0]))
		for x := 0; x < len(board[0]); x++ {
			fmt.Printf("%c ", board[y][x])
			key := board[y][x]
			dic[key] = append(dic[key], []int{y, x})
		}
		fmt.Println()
	}

	//for i := 0; i < len(markedExist); i++ {
	//	fmt.Println(markedExist[i])
	//}
	//
	//for key, value := range dic {
	//	fmt.Printf("%c : ", key)
	//	fmt.Println(value)
	//}

	res := false

	next := word[0]
	_, ok := dic[next]
	if !ok {
		return false
	}

	for i := 0; i < len(dic[next]); i++ {
		resetMarkedExist()
		res = res || dfsExist(board, word, 0, dic[next][i][0], dic[next][i][1])
	}

	fmt.Println(res)

	return res
}

func dfsExist(board [][]byte, word string, next int, y, x int) bool {
	if y == len(board) || x == len(board[0]) || y < 0 || x < 0 || word[next] != board[y][x] {
		return false
	}

	if markedExist[y][x] {
		return false
	}

	markedExist[y][x] = true

	fmt.Printf("next : %c y : %d x : %d\n", word[next], y, x)

	if next == len(word)-1 {
		return word[next] == board[y][x]
	}

	if dfsExist(board, word, next+1, y-1, x) ||
		dfsExist(board, word, next+1, y+1, x) ||
		dfsExist(board, word, next+1, y, x-1) ||
		dfsExist(board, word, next+1, y, x+1) {
		return true
	}

	markedExist[y][x] = false

	return false
}
