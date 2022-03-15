package main

var res [][]int

func subsets(nums []int) [][]int {
	res = make([][]int, 0)
	dfs(nums, 0, nil)
	return res
}

func dfs(nums []int, path int, prefix []int) {
	if path >= len(nums) {
		dst := make([]int, len(prefix))
		copy(dst, prefix)
		res = append(res, dst)
		return
	}
	dfs(nums, path+1, append(prefix, nums[path]))
	dfs(nums, path+1, prefix)
}

func subsets1(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	var bfs func(int)
	bfs = func(start int) {
		if start > len(nums) {
			return
		}
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			bfs(i + 1)
			path = path[:len(path)-1]
		}
	}
	bfs(0)
	return res
}

func main() {

}
