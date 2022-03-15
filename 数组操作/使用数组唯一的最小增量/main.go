package main

import "sort"

func minIncrementForUnique(nums []int) int {
	sort.Ints(nums)
	move := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			prev := nums[i]
			nums[i] = nums[i-1] + 1
			move += nums[i] - prev
		}
	}
	return move
}

func main() {

}
