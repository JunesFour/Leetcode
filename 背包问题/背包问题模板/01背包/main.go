package main

import "fmt"

// 二维动态规划
func backpack(w, v []int, cap int) int {
	f := make([][]int, len(w))
	for i := 0; i < len(w); i++ {
		f[i] = make([]int, cap+1)
	}

	//最开始放入第一件物品的情况
	for i := w[0]; i <= cap; i++ {
		f[0][i] = v[0]
	}

	for i := 1; i < len(w); i++ {
		for j := 0; j <= cap; j++ {
			if j < w[i] {
				f[i][j] = f[i-1][j]
			} else {
				f[i][j] = max(f[i-1][j], f[i][j-w[i]]+v[i])
			}
		}
	}
	return f[len(w)-1][cap]
}

// 一维动态规划
func backpack1(w, v []int, cap int) int {
	f := make([]int, cap+1)
	for i := 0; i < len(w); i++ {
		// 从大到小遍历时，当前的物品w[i]只有可能被选择一遍
		for j := cap; j >= w[i]; j-- {
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
	fmt.Println(backpack1([]int{1, 3, 4}, []int{15, 20, 30}, 4))
}
