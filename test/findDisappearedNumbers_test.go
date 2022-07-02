package main

import "testing"

func findDisappearedNumbers(nums []int) (res []int) {
	i := 0
	for i < len(nums) {
		// 如果当前元素nums[i]正好出现在他应该出现的地方
		if nums[i] == i+1 {
			i++
			continue
		}
		// 如果遍历到的元素，它应该出现的位置上已经有了一个与他一摸一样的数，说明重复了
		// 应该出现位置的数
		itemPos := nums[i] - 1
		if nums[i] == nums[itemPos] {
			i++
			continue
		}
		// 否则就交换
		nums[i], nums[itemPos] = nums[itemPos], nums[i]
	}
	for i, _ := range nums {
		if nums[i] != i+1 {
			res = append(res, i+1)
		}
	}
	return
}

func TestFindDisappearedNumbers(t *testing.T) {

}
