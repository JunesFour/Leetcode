package main

import (
	"fmt"
)

func productExceptSelf(nums []int) []int {
	if len(nums) == 2 {
		return []int{nums[1], nums[0]}
	}
	res := make([]int, len(nums))
	// 将res数组构建成从右到左的前缀和数组
	res[len(res)-1] = nums[len(nums)-1]
	for i := len(res) - 2; i >= 0; i-- {
		res[i] = nums[i] * res[i+1]
	}
	// 将nums数组构建成从左到右的前缀和数组
	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i] * nums[i-1]
	}
	// 第一个元素实际上就是顺序前缀和数组的第二个元素
	res[0] = res[1]
	// 备份res数组的最后一个元素
	end := res[len(res)-1]
	// 最后一个元素实际上就是逆序前缀和数组的倒数第二个元素
	res[len(res)-1] = nums[len(nums)-2]
	for i := 1; i <= len(res)-2; i++ {
		if i != len(res)-2 {
			res[i] = nums[i-1] * res[i+1]
		} else {
			res[i] = nums[i-1] * end
		}
	}
	return res
}

func productExceptSelf1(nums []int) []int {
	res := make([]int, len(nums))
	res[0] = 1
	// res[i]表示i左边的所有数字相乘
	for i := 1; i < len(res); i++ {
		res[i] = res[i-1] * nums[i-1]
	}
	r := 1
	// r 表示i右边的所有数字相乘
	for i := len(res) - 1; i >= 0; i-- {
		res[i] = res[i] * r
		r *= nums[i]
	}
	return res
}

func main() {
	fmt.Println(productExceptSelf1([]int{1, 2, 3, 4}))
}
