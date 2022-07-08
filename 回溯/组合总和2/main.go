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

// 实现2

func combinationSum21(candidates []int, target int) (res [][]int) {
	var comb []int
	// 因为不能重复，所以需要去重
	sort.Ints(candidates)
	var dfs func(target int, index int)
	dfs = func(target int, index int) {
		if target == 0 {
			res = append(res, append([]int{}, comb...))
			return
		}
		// 1. 在循环里面去重
		// 2. 在循环里面把dfs(target, index + 1)这一步已经默认执行过了
		for i := index; i < len(candidates); i++ {
			// 去重
			if i != index && candidates[i] == candidates[i-1] {
				continue
			}
			// 剪枝
			if target < candidates[i] {
				return
			}
			comb = append(comb, candidates[i])
			dfs(target-candidates[i], i+1)
			comb = comb[:len(comb)-1]
		}
	}
	dfs(target, 0)
	return
}

func main() {

}
