package main

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var res [][]int
	dfs(candidates, nil, target, 0, &res)
	return res
}

func dfs(candidates []int, nums []int, target int, left int, res *[][]int) {
	// 递归出口
	if target == 0 {
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*res = append(*res, tmp)
		return
	}
	for i := left; i < len(candidates); i++ {
		if i != left && candidates[i] == candidates[i-1] { // 去重
			continue
		}
		if target < candidates[i] { // 剪枝
			return
		}
		dfs(candidates, append(nums, candidates[i]), target-candidates[i], i+1, res)
	}
}

func main() {

}
