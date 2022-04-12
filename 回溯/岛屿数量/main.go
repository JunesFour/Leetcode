package main

import "fmt"

func numIslands(grid [][]byte) int {
	var res int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				res++
				dfs(i, j, grid)
			}
		}
	}
	return res
}
func dfs(i, j int, grid [][]byte) {
	// 如果出界，就返回
	if i > len(grid)-1 || i < 0 || j > len(grid[0])-1 || j < 0 {
		return
	}
	// 如果遇到0，就返回
	if grid[i][j] == '0' {
		return
	}
	// 将原来是1的位置标记为2
	grid[i][j] = '0'
	// 向四个方向递归
	dfs(i+1, j, grid)
	dfs(i-1, j, grid)
	dfs(i, j-1, grid)
	dfs(i, j+1, grid)
}

func main() {
	fmt.Println(numIslands([][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}}))
}
