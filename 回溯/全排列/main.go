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

// 使用标记数组
func permute1(nums []int) (res [][]int) {
	var cur []int
	visited := make([]bool, len(nums))
	var dfs func(index int)
	dfs = func(index int) {
		if len(cur) == len(nums) {
			res = append(res, append([]int{}, cur...))
			return
		}

		for i := 0; i < len(nums); i++ {
			if visited[i] {
				continue
			}
			visited[i] = true
			cur = append(cur, nums[i])
			dfs(i)
			visited[i] = false
			cur = cur[:len(cur)-1]
		}
	}
	dfs(0)
	return res
}

func main() {

}
