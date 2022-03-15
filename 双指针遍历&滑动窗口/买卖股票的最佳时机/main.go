package main

import "math"

func maxProfit(prices []int) int {
	minValue := math.MaxInt64
	maxValue := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minValue {
			minValue = prices[i]
		} else if prices[i]-minValue > maxValue {
			maxValue = prices[i] - minValue
		}
	}
	return maxValue
}

func main() {

}
