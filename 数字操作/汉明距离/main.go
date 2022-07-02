package main

import (
	"fmt"
	"math/bits"
)

func hammingDistance(x int, y int) int {
	return bits.OnesCount(uint(x ^ y))
}

func main() {
	fmt.Println(hammingDistance(3, 13))
}
