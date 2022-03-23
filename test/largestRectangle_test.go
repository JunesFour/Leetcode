package main

import (
	"fmt"
	"testing"
)

// 滑动窗口
func largestRectangleArea(heights []int) int {
	minHeight, maxArea := heights[0], 0
	var UpdateArea func(i int)
	UpdateArea = func(i int) {
		if heights[i] >= minHeight {
			maxArea += minHeight
		} else {
			minHeight = heights[i]
			maxArea = heights[i]
		}
	}
	for i := 0; i < len(heights); i++ {
		UpdateArea(i)
	}
	return maxArea
}

func TestLargestRectangle(t *testing.T) {
	fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
}
