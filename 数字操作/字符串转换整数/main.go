package main

import (
	"fmt"
	"math"
)

func myAtoi(s string) int {
	i, num, isNegative, n := 0, 0, 1, len(s)
	for i < n && s[i] == ' ' {
		i++
	}
	if i < n {
		if s[i] == '-' {
			isNegative = -1
			i++
		} else if s[i] == '+' {
			i++
		}
	}
	for i < n && s[i] >= '0' && s[i] <= '9' {
		num = num*10 + int(s[i]-'0')
		if isNegative*num < (-math.MaxInt32 - 1) {
			return (-math.MaxInt32 - 1)
		}
		if isNegative*num > math.MaxInt32 {
			return math.MaxInt32
		}
		i++
	}
	return isNegative * num
}

func main() {
	fmt.Println(myAtoi("-24"))
}
