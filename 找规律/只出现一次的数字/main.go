package main

import "sort"

func singleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	sort.Ints(nums)
	for i, j := 0, 1; i < len(nums); {
		if i == len(nums)-1 {
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

// 使用异或运算
func singleNumber1(nums []int) int {
	var res int
	for _, v := range nums {
		res ^= v
	}
	return res
}

func main() {

}
