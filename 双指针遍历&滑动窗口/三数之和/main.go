package main

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var saveArray [][]int
	var addnum int
	//index为中间值
	var index, start, end int
	for index = 1; index < len(nums)-1; index++ {
		start, end = 0, len(nums)-1
		//去重的关键步骤
		if index > 1 && nums[index] == nums[index-1] {
			start = index - 1
		}
		for start < index && index < end {
			addnum = nums[start] + nums[index] + nums[end]
			if start > 0 && nums[start] == nums[start-1] {
				start++
				continue
			}
			if end < len(nums)-1 && nums[end] == nums[end+1] {
				end--
				continue
			}
			if addnum == 0 {
				saveArray = append(saveArray, []int{nums[start], nums[index], nums[end]})
				start++
				end--
			} else if addnum > 0 {
				end--
			} else {
				start++
			}
		}
	}
	return saveArray
}

func main() {

}
