package algirithm

import (
	"leetcodetop100/kit"
)

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	min := prices[0]
	max := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else if max < prices[i]-min {
			max = prices[i] - min
		}
	}
	return max
}

func maxProfitI(prices []int) int {

	n := len(prices)
	if n <= 1 {
		return 0
	}

	dp_i_0 := 0
	dp_i_1 := -1 << 31

	for i := 0; i < n; i++ {
		dp_i_0 = kit.Max(dp_i_0, dp_i_1+prices[i])
		dp_i_1 = kit.Max(dp_i_1, -prices[i])
	}

	return dp_i_0
}
