package main

import "math"

func numSquares(n int) int {
	//定义
	dp := make([]int, n+1)
	// 初始化
	dp[0] = 0
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt32
	}
	// 遍历物品
	for i := 1; i <= n; i++ {
		// 遍历背包
		for j := i * i; j <= n; j++ {
			dp[j] = min(dp[j], dp[j-i*i]+1)
		}
	}

	return dp[n]
}

// 版本二,先遍历背包, 再遍历物品
func numSquares2(n int) int {
	//定义
	dp := make([]int, n+1)
	// 初始化
	dp[0] = 0
	// 遍历背包
	for j := 1; j <= n; j++ {
		//初始化
		dp[j] = math.MaxInt32
		// 遍历物品
		for i := 1; i <= n; i++ {
			if j >= i*i {
				dp[j] = min(dp[j], dp[j-i*i]+1)
			}
		}
	}

	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
