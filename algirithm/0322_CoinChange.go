package algirithm

func CoinChange(coins []int, amount int) int {
	if coins == nil || len(coins) == 0 || amount < 0 {
		return -1
	}

	var min = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	dp := make([]int, amount+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = 1 << 31
	}
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if i >= coins[j] {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}

	return dp[amount]
}
