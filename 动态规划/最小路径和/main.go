package main

import "fmt"

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	dp[0] = make([]int, n)
	dp[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = grid[i][0]
		dp[i][0] += dp[i-1][0]
	}
	for j := 1; j < n; j++ {
		dp[0][j] = grid[0][j]
		dp[0][j] += dp[0][j-1]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = grid[i][j]
			dp[i][j] += min(dp[i-1][j], dp[i][j-1])
		}
	}
	return dp[m-1][n-1]
}
func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {
	fmt.Println(minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
}
