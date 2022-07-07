package main

import "fmt"

func maxCoins(nums []int) int {
	// 新开辟一个下标从0-n+1的数组
	points := make([]int, len(nums)+2)
	// 下标0和n+1作为边界，值都为1
	points[0], points[len(nums)+1] = 1, 1
	for i := 1; i <= len(nums); i++ {
		points[i] = nums[i-1]
	}

	// dp[i][j]表示从i到j区间（开区间）之间的最优解
	dp := make([][]int, len(nums)+2)
	for i := 0; i < len(nums)+2; i++ {
		dp[i] = make([]int, len(nums)+2)
	}

	// 状态转移
	// i从下往上
	for i := len(nums); i >= 0; i-- {
		// j从左往右
		for j := i + 1; j < len(nums)+2; j++ {
			// k表示最后被戳破的气球
			for k := i + 1; k < j; k++ {
				dp[i][j] = max(dp[i][j], dp[i][k]+dp[k][j]+points[i]*points[j]*points[k])
			}
		}
	}

	return dp[0][len(nums)+1]
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	fmt.Println(maxCoins([]int{3, 1, 5, 8}))
}
