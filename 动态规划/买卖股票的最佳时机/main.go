package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	minValue := math.MaxInt64
	maxValue := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minValue {
			minValue = prices[i]
		} else if prices[i]-minValue > maxValue {
			maxValue = prices[i] - minValue
		}
	}
	return maxValue
}

func maxProfit1(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	// dp[i][0]表示今天不持股 dp[i][1]表示今天持股
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2)
	}
	// 第一天不持股
	dp[0][0] = 0
	// 第一天持股
	dp[0][1] = -prices[0]
	// 从第2天开始
	for i := 1; i < len(prices); i++ {
		// 今天不持股
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		// 今天持股
		dp[i][1] = max(dp[i-1][1], -prices[i])
	}
	return dp[len(prices)-1][0]
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
