package main

import "math"

// 双指针
func maxArea(height []int) (Max int) {
	left, right := 0, len(height)-1
	Max = (right - left) * int(math.Min(float64(height[left]), float64(height[right])))
	for left < right {
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
		area := (right - left) * int(math.Min(float64(height[left]), float64(height[right])))
		Max = int(math.Max(float64(Max), float64(area)))
	}
	return
}

func maxArea1(height []int) (Max int) {
	left, right := 0, len(height)-1
	Max = (right - left) * min(height[left], height[right])
	for left < right {
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
		area := (right - left) * min(height[left], height[right])
		Max = max(Max, area)
	}
	return
}
func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {

}
