package 每日1题_3月

func maxProfit(prices []int) int {
	if prices == nil || len(prices) == 0 {
		return 0
	}

	dp_i_0 := 0
	dp_i_1 := -1 << 31

	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := 0; i < len(prices); i++ {
		dp_i_0 = max(dp_i_0, dp_i_1+prices[i])
		dp_i_1 = max(dp_i_1, -prices[i])
	}

	return dp_i_0
}
