package main

func search(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	if n == 1 {
		if target == nums[0] {
			return 0
		} else {
			return -1
		}
	}
	l, r := 0, n-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}
		// 如果左边有序
		if nums[0] <= nums[mid] {
			// 如果target恰好在左边区间
			if nums[0] <= target && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
			// 如果右边有序
		} else {
			// 如果target恰好在右边区间
			if nums[mid] < target && target <= nums[n-1] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}

func main() {

}
