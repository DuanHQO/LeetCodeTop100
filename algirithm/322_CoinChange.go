package algirithm

import (
	"fmt"
	"leetcodetop100/kit"
)

func CoinChange(coins []int, amount int) int {
	if coins == nil || len(coins) == 0 || amount < 0 {
		return -1
	}

	//sort.Ints(coins)
	//
	////fmt.Println(coins)
	//
	//var result [][]int
	//min := 1 << 31
	//
	//var dfs func(int, int, []int)
	//dfs = func(start int, left int, pre []int) {
	//	//fmt.Printf("start %d left %d pre %v\n", start, left, pre)
	//	if left < 0 {
	//		return
	//	}
	//
	//	if left == 0 {
	//
	//		if len(pre) <= min {
	//			min = len(pre)
	//			temp := make([]int, len(pre))
	//			for i := 0; i < len(temp); i++ {
	//				temp[i] = pre[i]
	//			}
	//			result = append(result, temp)
	//		}
	//		return
	//	}
	//
	//	for i := start; i >= 0 && left - coins[i] >= 0 && len(pre) <= min; i-- {
	//		pre = append(pre, coins[i])
	//		//fmt.Println(pre)
	//		dfs(start, left - coins[i], pre)
	//		pre = pre[:len(pre) - 1]
	//	}
	//}
	//
	//start := len(coins) - 1
	//for i := len(coins) - 1; i >= 0; i-- {
	//	if coins[i] <= amount {
	//		start = i
	//		break
	//	}
	//}
	//
	//dfs(start, amount, []int{})
	//
	//if result == nil || len(result) == 0 {
	//	min = -1
	//}
	//
	//fmt.Println(min)

	dp := make([]int, amount+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = 1 << 31
	}
	fmt.Println(dp)
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if i >= coins[j] {
				fmt.Printf("i %d coins[j] %d\n", i, coins[j])
				dp[i] = kit.Min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}
	fmt.Println(dp)
	if dp[amount] > amount {
		return -1
	} else {
		return dp[amount]
	}

	//return min
}
