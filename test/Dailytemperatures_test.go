package main

import (
	"fmt"
	"testing"
)

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	res := make([]int, n)
	var stack []int
	for i := 0; i < n; i++ {
		temperature := temperatures[i]
		// 为栈中的元素赋值天数
		for len(stack) > 0 && temperature > temperatures[stack[len(stack)-1]] {
			// 给栈顶元素赋值天数
			res[stack[len(stack)-1]] = i - stack[len(stack)-1]
			// 弹出栈顶元素
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return res
}

func TestDailytemperatures(t *testing.T) {
	//  1  0  0   2 1  0  0   0  0  0
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}
