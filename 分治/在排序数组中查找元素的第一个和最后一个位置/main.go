package main

func searchRange(nums []int, target int) []int {
	l, r, midIndex, n := 0, len(nums)-1, -1, len(nums)
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			midIndex = mid
			break
		}
		if nums[mid] < target && target <= nums[n-1] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	if midIndex == -1 {
		return []int{-1, -1}
	}

	start, end := 0, 0
	for i := midIndex; i >= 0; i-- {
		if nums[i] == target {
			start = i
		} else {
			break
		}
	}
	for i := midIndex; i <= n-1; i++ {
		if nums[i] == target {
			end = i
		} else {
			break
		}
	}
	return []int{start, end}
}

func main() {

}
