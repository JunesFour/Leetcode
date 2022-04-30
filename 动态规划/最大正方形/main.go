package main

func maximalSquare(matrix [][]byte) int {
	maxSide := 0
	dp := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[i]))
		for j := 0; j < len(matrix[i]); j++ {
			dp[i][j] = int(matrix[i][j] - '0')
			if dp[i][j] == 1 {
				maxSide = 1
			}
		}
	}
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if dp[i][j] == 1 {
				dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
				if dp[i][j] > maxSide {
					maxSide = dp[i][j]
				}
			}

		}
	}
	return maxSide * maxSide
}

func min(a, b, c int) int {
	return minTwo(minTwo(a, b), c)
}
func minTwo(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {

}
