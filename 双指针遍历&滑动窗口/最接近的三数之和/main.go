package main

import "sort"

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n, sumClosest := len(nums), nums[0]+nums[1]+nums[2]
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		j, k := i+1, n-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == target {
				return target
			}
			if abs(sum-target) < abs(sumClosest-target) {
				sumClosest = sum
			}
			if sum > target {
				k--
			} else {
				j++
			}
		}
	}
	return sumClosest
}

func abs(num int) int {
	if num < 0 {
		num = -num
	}
	return num
}

func main() {

}
