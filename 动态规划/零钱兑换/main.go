package main

import (
	"fmt"
	"math"
)

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i, _ := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			if dp[j-coins[i]] != math.MaxInt32 {
				dp[j] = min(dp[j-coins[i]]+1, dp[j])
			}
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}
func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
}
