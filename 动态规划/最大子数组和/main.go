package main

import (
	"fmt"
)

func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	maxSum := nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = nums[i]
		if dp[i-1]+dp[i] > dp[i] {
			dp[i] += dp[i-1]
		}
		if dp[i] > maxSum {
			maxSum = dp[i]
		}
	}
	return maxSum
}

func main() {
	fmt.Println(maxSubArray([]int{5, 4, -1, 7, 8}))
}
