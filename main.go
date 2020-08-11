package main

import (
	_ "fmt"
	_ "github.com/chai2010/errors"
)
import _ "github.com/nosixtools/timewheel"
import _ "github.com/rfyiamcool/go-timewheel"
import "leetcodetop100/algirithm"

func main() {
	//每日1题3月.HasGroupsSizeX([]int{1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3})

	//fmt.Println(algirithm.CountBinarySubstrings("00110011"))
	algirithm.Solve([][]byte{
		{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'},
	})
}
