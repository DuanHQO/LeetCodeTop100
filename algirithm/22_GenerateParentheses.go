package algirithm

func GenerateParenthesis(n int) []string {
	var res []string

	if n == 0 {
		return res
	}

	dfs22(n, n, "", &res)

	return res
}

func dfs22(left int, right int, curr string, res *[]string) {
	if left == 0 && right == 0 {
		*res = append(*res, curr)
		return
	}

	if left > 0 {
		dfs22(left-1, right, curr+"(", res)
	}

	if right > 0 && left < right {
		dfs22(left, right-1, curr+")", res)
	}
}
