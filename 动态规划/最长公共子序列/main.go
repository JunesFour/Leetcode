package main

import "fmt"

func longestCommonSubsequence(text1 string, text2 string) int {
	len1, len2 := len(text1), len(text2)
	dp := make([][]int, len1+1)
	for i := 0; i <= len1; i++ {
		dp[i] = make([]int, len2+1)
	}
	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			if text1[i-1] == text2[j-1] {
				dp[i][j] = max(dp[i][j], dp[i-1][j-1]+1)
			}
		}
	}
	return dp[len1][len2]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	fmt.Println(longestCommonSubsequence("abcde", "ace"))
}
