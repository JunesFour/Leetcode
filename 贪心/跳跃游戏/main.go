package main

import "fmt"

func canJump(nums []int) bool {
	// 如果数组长度为1，直接返回true
	if len(nums) < 2 {
		return true
	}
	maxRange := 0
	for i := 0; i < len(nums)-1; i++ {
		// 更新能覆盖到的最远值
		if i+nums[i] > maxRange {
			maxRange = i + nums[i]
		}
		// 如果某一个元素为0 且 能覆盖的最大范围不超过这个元素的位置，直接返回false
		if nums[i] == 0 && maxRange <= i {
			return false
		}
		// 覆盖范围超过最后一个元素的位置，直接返回true
		if maxRange >= len(nums)-1 {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(canJump([]int{0, 2, 3}))
}
