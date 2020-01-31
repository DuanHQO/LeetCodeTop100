package algirithm

func GenerateParenthesis(n int) []string {
	var res []string

	if n == 0 {
		return res
	}

	dfs(n, n, "", &res)

	return res
}

func dfs(left int, right int, curr string, res *[]string) {
	if left == 0 && right == 0 {
		*res = append(*res, curr)
		return
	}

	if left > 0 {
		dfs(left-1, right, curr+"(", res)
	}

	if right > 0 && left < right {
		dfs(left, right-1, curr+")", res)
	}
}
