package main

import "fmt"

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	step1, step2, step3 := 1, 2, 0
	for i := 3; i <= n; i++ {
		step3 = step1 + step2
		step1 = step2
		step2 = step3
	}
	return step3
}

func main() {
	fmt.Println(climbStairs(5))
}
