package main

import "fmt"

func isHappy(n int) bool {
	if n < 2 {
		return true
	}
	slow, fast := n, sumOfSquares(n)
	for slow != fast {
		if sumOfSquares(fast) == 1 {
			return true
		}
		slow = sumOfSquares(slow)
		fast = sumOfSquares(sumOfSquares(fast))
	}
	return false
}

func sumOfSquares(num int) (sum int) {
	for num != 0 {
		sum += (num % 10) * (num % 10)
		num /= 10
	}
	return sum
}

func main() {
	fmt.Println(isHappy(19))
}
