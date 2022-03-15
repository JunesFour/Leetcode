package main

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slow, fast := 0, 1
	for fast < len(nums) {
		if nums[slow] != nums[fast] {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	return slow + 1
}

func removeDuplicates1(nums []int) int {
	slow := 0
	p := -10001
	for i := range nums {
		if nums[i] > p {
			p = nums[i]
			nums[slow] = nums[i]
			slow++
		}
	}
	return slow
}

func main() {

}
