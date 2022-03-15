package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	num := 0
	tmp := x
	for tmp != 0 {
		num = num*10 + tmp%10
		tmp /= 10
	}
	return num == x
}

func main() {
	fmt.Println(isPalindrome(-121))
}
