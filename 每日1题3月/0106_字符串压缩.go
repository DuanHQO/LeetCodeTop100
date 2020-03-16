package 每日1题3月

import (
	"fmt"
	"strconv"
)

func CompressString(S string) string {
	if S == "" {
		return ""
	}

	fmt.Println(S)
	fmt.Println(len(S))

	ret := ""
	cur := 1
	curs := S[0:1]

	for i := 1; i < len(S); i++ {
		if S[i] == S[i-1] {
			cur++
			continue
		}
		ret = ret + curs + strconv.Itoa(cur)
		curs = S[i : i+1]
		cur = 1
	}
	ret = ret + curs + strconv.Itoa(cur)

	fmt.Println(ret)
	fmt.Println(len(ret))

	if len(S) <= len(ret) {
		return S
	}

	return ret
}
