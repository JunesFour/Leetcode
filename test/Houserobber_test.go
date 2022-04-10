package main

import (
	"fmt"
	"math"
	"testing"
)

// 1.加上当前 跳2 跳3
// 2.不加上当前 跳1 跳2
// 4 3 2 3
func rob(nums []int) int {
	Max := math.MinInt32
	var dfs func(res int, i int)
	dfs = func(res int, i int) {
		// 处理边界情况
		if i > len(nums)-1 {
			fmt.Println("res: ", res)
			if res > Max {
				Max = res
			}
			return
		}
		if i == len(nums)-1 {
			res += nums[i]
			fmt.Println("res: ", res)
			if res > Max {
				Max = res
			}
			return
		}
		if i == len(nums)-2 {
			if res+nums[i] > res+nums[i+1] {
				res += nums[i]
			} else {
				res += nums[i+1]
			}
			fmt.Println("res: ", res)
			if res > Max {
				Max = res
			}
			return
		}
		if i == len(nums)-3 {
			if res+nums[i]+nums[i+2] > res+nums[i+1] {
				res += nums[i] + nums[i+2]
			} else {
				res += nums[i+1]
			}
			fmt.Println("res: ", res)
			if res > Max {
				Max = res
			}
			return
		}

		// 往下递归
		dfs(res, i+1)
		dfs(res, i+2)
		dfs(res+nums[i], i+2)
		dfs(res+nums[i], i+3)
		// 0 1 2 3
	}
	dfs(0, 0)
	return Max
}

// 动态规划
// dp[i] = Max(dp[i - 1], dp[i] + dp[i - 2], dp[i] + dp[i - 3], dp[i - 1] + dp[i - 3])
// 1,2,3,4
func rob1(nums []int) int {
	// 特判
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(dp[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[len(nums)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func TestHouseRobber(t *testing.T) {
	fmt.Println(rob1([]int{2, 7, 9, 3, 1}))
}
