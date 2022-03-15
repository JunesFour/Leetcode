package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	num := 0
	for x != 0 {
		num = num*10 + x%10
		x /= 10
		if num < -math.MaxInt32 || num > math.MaxInt32 {
			return 0
		}
	}
	return num
}

func main() {
	fmt.Println(reverse(-120))
}
