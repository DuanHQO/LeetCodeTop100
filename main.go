package main

import (
	"fmt"
	_ "fmt"
	_ "github.com/chai2010/errors"
	"leetcodetop100/kit"
)
import _ "github.com/nosixtools/timewheel"
import _ "github.com/rfyiamcool/go-timewheel"

//import "leetcodetop100/algirithm"

func main() {
	//每日1题3月.HasGroupsSizeX([]int{1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3})

	//fmt.Println(algirithm.CountBinarySubstrings("00110011"))
	//algirithm.Solve([][]byte{
	//	{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'},
	//})

	arr := []int{12, 5788, 6, 515, 48, 656, 87, 2323, 665, 7, 4, -548}
	kit.MergeSort(arr, 0, 11)
	fmt.Println(arr)
}
