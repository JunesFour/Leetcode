package main

func majorityElement(nums []int) int {
	// 摩尔投票法
	res, count := -1, 0
	for _, v := range nums {
		if count == 0 {
			res = v
		}
		if v == res {
			count++
		} else {
			count--
		}
	}
	return res
}

func main() {

}
