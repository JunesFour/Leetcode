package main

import "fmt"

func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(word2)+1)
	}
	for i := 1; i < len(dp); i++ {
		dp[i][0] = i
	}
	for j := 1; j < len(dp[0]); j++ {
		dp[0][j] = j
	}
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			// 当word1[i-1]不等于word2[j - 1]时
			dp[i][j] = min(min(dp[i][j-1], dp[i-1][j]), dp[i-1][j-1]) + 1
			// 当word1[i-1]等于word2[j - 1]时
			if word1[i-1] == word2[j-1] {
				dp[i][j] = min(dp[i][j], dp[i-1][j-1])
			}
		}
	}
	return dp[len(dp)-1][len(dp[0])-1]
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {
	fmt.Println(minDistance("horse", "ros"))
}
