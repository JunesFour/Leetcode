package main

import "fmt"

func countSubstrings(s string) int {
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	res := 0
	for i := 0; i < len(s); i++ {
		for j := 0; j <= i; j++ {
			if s[i] == s[j] && (i-j < 2 || dp[j+1][i-1]) {
				dp[j][i] = true
				res++
			}
		}
	}
	return res
}

func main() {
	fmt.Println()
}
