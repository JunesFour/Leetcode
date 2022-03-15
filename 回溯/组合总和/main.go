package main

import "fmt"

func combinationSum(candidates []int, target int) (res [][]int) {
	var comb []int
	var dfs func(target int, index int)
	dfs = func(target int, index int) {
		if index == len(candidates) {
			return
		}
		if target == 0 {
			res = append(res, append([]int{}, comb...))
			/*tmp := make([]int, len(comb))
			copy(tmp, comb)
			res = append(res, tmp)*/
			return
		}
		// 直接跳过
		dfs(target, index+1)
		// 减去当前的元素
		if target-candidates[index] >= 0 {
			comb = append(comb, candidates[index])
			dfs(target-candidates[index], index)
			comb = comb[:len(comb)-1]
		}
	}
	dfs(target, 0)
	return
}

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
}
