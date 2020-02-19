package algirithm

import "leetcodetop100/kit"

func maxProfitII(prices []int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	}
	dp_i_0 := 0
	dp_i_1 := -1 << 31
	for i := 0; i < n; i++ {
		temp := dp_i_0
		dp_i_0 = kit.Max(dp_i_0, dp_i_1+prices[i])
		dp_i_1 = kit.Max(dp_i_1, temp-prices[i])
	}
	return dp_i_0
}
