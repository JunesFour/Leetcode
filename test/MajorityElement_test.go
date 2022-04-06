package main

import (
	"fmt"
	"testing"
)

func majorityElement(nums []int) int {
	// 摩尔投票法
	res, count := -1, 0
	for _, v := range nums {
		if count == 0 {
			res = v
		}
		if v == res {
			count++
		} else {
			count--
		}
	}
	return res
}

func TestMajorityElement(t *testing.T) {
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}
