package main

import "fmt"

func backpack(v, w []int, cap int) int {
	f := make([]int, cap+1, cap+1)
	for i := 0; i < len(w); i++ {
		// 从小到大遍历时，可以使每个物体被多次选择
		for j := w[i]; j <= cap; j++ {
			f[j] = max(f[j], f[j-w[i]]+v[i])
		}
	}
	return f[cap]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	fmt.Println(backpack([]int{5, 7, 3}, []int{1, 2, 1}, 5))
}
