package main

func leastInterval(tasks []byte, n int) int {
	var nums [26]int
	maxCount := 0
	// 统计每个字母的出现次数，并找出出现次数最多的字母
	for _, task := range tasks {
		nums[task-'A']++
		maxCount = max(nums[task-'A'], maxCount)
	}
	res := (maxCount - 1) * (n + 1)
	for i := 0; i < 26; i++ {
		if maxCount == nums[i] {
			res++
		}
	}
	return max(res, len(tasks))
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {

}
