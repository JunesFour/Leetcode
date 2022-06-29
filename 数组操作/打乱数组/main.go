package main

import "math/rand"

type Solution struct {
	nums   []int
	reNums []int
}

func Constructor(nums []int) Solution {
	return Solution{nums: nums, reNums: append([]int{}, nums...)}
}

func (this *Solution) Reset() []int {
	return this.reNums
}

func (this *Solution) Shuffle() []int {
	for n := len(this.nums); n > 0; n-- {
		randIndex := rand.Intn(n)
		this.nums[n-1], this.nums[randIndex] = this.nums[randIndex], this.nums[n-1]
	}
	return this.nums
}

func main() {

}
