package main

func findUnsortedSubarray(nums []int) int {
	n := len(nums)
	min, max := nums[n-1], nums[0]
	begin, end := -1, -1
	for i := 0; i < n; i++ {
		if nums[i] >= max {
			max = nums[i]
		} else {
			end = i
		}
		if nums[n-i-1] <= min {
			min = nums[n-i-1]
		} else {
			begin = n - i - 1
		}
	}
	// 特判
	if end == -1 {
		return 0
	}
	return end - begin + 1
}

func main() {

}
