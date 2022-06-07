package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	n := len(prices)
	dp := make([][3]int, n)
	dp[0][0] = -prices[0]
	for i := 1; i < n; i++ {
		// 第i天持有股票，以往就持有或者前一天买的
		dp[i][0] = max(dp[i-1][0], dp[i-1][2]-prices[i])
		// 第i天不持有股票且在冷冻期，只能是前一天卖的
		dp[i][1] = dp[i-1][0] + prices[i]
		// 第i天不持有股票且不在冷冻期，继承前一天的两种状态
		dp[i][2] = max(dp[i-1][1], dp[i-1][2])
	}
	return max(dp[n-1][1], dp[n-1][2])
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	fmt.Println(maxProfit([]int{1, 2, 3, 0, 2}))
}
