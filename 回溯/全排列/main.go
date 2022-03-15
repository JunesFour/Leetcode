package main

func permute(nums []int) [][]int {
	var res [][]int
	dfs(nums, []int{}, &res)
	return res
}

func dfs(nums []int, cur []int, res *[][]int) {
	if len(nums) == len(cur) {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*res = append(*res, tmp)
		return
	}
	for i := 0; i < len(nums); i++ {
		exist := false
		for j := 0; j < len(cur); j++ {
			if nums[i] == cur[j] {
				exist = true
				break
			}
		}
		if exist == true {
			continue
		}
		dfs(nums, append(cur, nums[i]), res)
	}
}

func main() {

}
