package main

import (
	"fmt"
	"testing"
)

func countBits(n int) []int {
	res := make([]int, n+1)
	for i := 0; i <= n; i++ {
		res[i] = countOne(i)
	}
	return res
}
func countOne(num int) (res int) {
	for num != 0 {
		if num%2 == 1 {
			res++
		}
		num /= 2
	}
	return
}
func TestCountBits(t *testing.T) {
	fmt.Println(countBits(5))
}
