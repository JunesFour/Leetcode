package main

import (
	"fmt"
	"sort"
	"testing"
)

func singleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0] // 1 1 2 2 3 4 4
	}
	sort.Ints(nums)
	for i, j := 0, 1; i < len(nums); {
		if i == len(nums)-1 { // 1 1 2 2 4
			return nums[i]
		}
		if nums[i] != nums[j] {
			return nums[i]
		}
		i += 2
		j += 2
	}
	return 1
}

func singleNumber1(nums []int) int {
	var res int
	for _, v := range nums {
		res ^= v
	}
	return res
}

func TestSingleNumber(t *testing.T) {
	fmt.Println(5 ^ 5)
	fmt.Println(5 ^ 6)
}
