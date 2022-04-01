package main

import (
	"fmt"
	"math"
)

func maxProduct(nums []int) int {
	// 维护一个当前的最大值和最小值，如果遇到负数，原来的最大值会变成最小值，最小值会变成最大值
	currMax, currMin, Max := 1, 1, math.MinInt
	for _, v := range nums {
		// 当前值为负数，最大值和最小值做一个交换
		// 因为做了交换，所以当前最大值一定大于当前最小值
		if v < 0 {
			currMax, currMin = currMin, currMax
		}
		// 更新当前最大值和最小值
		currMax = max(v*currMax, v)
		currMin = min(v*currMin, v)

		// 更新最大值
		Max = max(Max, currMax)
	}
	return Max
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {
	fmt.Println(maxProduct([]int{-1, 0, -2}))
}
