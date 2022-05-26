package main

import "fmt"

// 价值、体积、数量
func backpack(v, w, s []int, cap int) int {
	f := make([]int, cap+1, cap+1)
	for i := 0; i < len(w); i++ {
		for j := cap; j >= w[i]; j-- {
			for k := 1; k <= s[i]; k++ {
				if j >= k*w[i] {
					f[j] = max(f[j], f[j-k*w[i]]+k*v[i])
				}
			}
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
	fmt.Println(backpack([]int{}, []int{}, []int{}, 5))
}
