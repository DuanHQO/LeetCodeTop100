package main

import (
	"fmt"
	_ "github.com/chai2010/errors"
)
import _ "github.com/nosixtools/timewheel"
import _ "github.com/rfyiamcool/go-timewheel"
import "leetcodetop100/algirithm"

func main() {
	//每日1题3月.HasGroupsSizeX([]int{1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3})
	res := algirithm.GenerateParenthesis(3)
	for _, re := range res {
		fmt.Println(re)
	}
}
