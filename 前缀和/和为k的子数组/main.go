package main

import "fmt"

func subarraySum(nums []int, k int) int {
	mp := make(map[int]int)
	mp[0] = 1
	pre, count := 0, 0
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if _, ok := mp[pre-k]; ok {
			count += mp[pre-k]
		}
		mp[pre] += 1
	}
	return count
}

func main() {
	fmt.Println(subarraySum([]int{5, 1, 2, 5}, 3))
}
