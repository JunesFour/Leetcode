package main

import "fmt"

// 超时
func largestRectangleArea(heights []int) int {
	area := 0
	for i := 0; i < len(heights); i++ {
		w := 1
		h := heights[i]
		j := i - 1
		for j >= 0 && heights[j] >= h {
			w++
			j--
		}
		j = i + 1
		for j < len(heights) && heights[j] >= h {
			w++
			j++
		}
		area = max(area, w*h)
	}
	return area
}

func largestRectangleArea1(heights []int) int {
	// 单调栈
	stack := make([]int, 0)
	// 添加一个哨兵节点
	stack = append(stack, -1)
	heights = append(heights, 0)
	res := 0
	for i := 0; i < len(heights); i++ {
		// 当当前元素比栈顶元素小的时候，栈顶元素就可以确定左右边界
		for len(stack) > 1 && heights[stack[len(stack)-1]] > heights[i] {
			top := stack[len(stack)-1]
			// 出栈
			stack = stack[:len(stack)-1]
			// 左边界，右边界就是当前元素
			left := stack[len(stack)-1]
			// 更新矩形最大面积
			res = max(res, (i-left-1)*heights[top])
		}
		stack = append(stack, i)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
}
