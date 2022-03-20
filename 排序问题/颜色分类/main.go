package main

import "fmt"

func sortColors(nums []int) {
	var red, white, blue int
	for _, v := range nums {
		if v == 0 {
			red++
		} else if v == 1 {
			white++
		} else if v == 2 {
			blue++
		}
	}
	var index int
	for red != 0 {
		nums[index] = 0
		index++
		red--
	}
	for white != 0 {
		nums[index] = 1
		index++
		white--
	}
	for blue != 0 {
		nums[index] = 2
		index++
		blue--
	}
}

func sortColors1(nums []int) {
	var p0, p1 int
	for i, v := range nums {
		if v == 0 {
			nums[i], nums[p0] = nums[p0], nums[i]
			if p0 < p1 {
				nums[i], nums[p1] = nums[p1], nums[i]
			}
			p0++
			p1++
		}
		if v == 1 {
			nums[i], nums[p1] = nums[p1], nums[i]
			p1++
		}
	}
}

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Println(nums)
}
