package main

import (
	"fmt"
	"testing"
)

func searchMatrix(matrix [][]int, target int) bool {
	var find bool
	visited := make([][]bool, len(matrix))
	for i := 0; i < len(matrix); i++ {
		visited[i] = make([]bool, len(matrix[i]))
	}
	var dfs func(i, j int)
	dfs = func(i, j int) {
		visited[i][j] = true
		// 找到数字
		if matrix[i][j] == target {
			find = true
			return
		}
		if find == true {
			return
		}
		//fmt.Println("i: ", i, ", j: ", j)
		// 未出下边界
		if i+1 < len(matrix) && matrix[i+1][j] <= target && !visited[i+1][j] {
			dfs(i+1, j)
		}
		// 未出下边界
		if j+1 < len(matrix[0]) && matrix[i][j+1] <= target && !visited[i][j+1] {
			dfs(i, j+1)
		}
	}
	dfs(0, 0)
	return find
}

func TestSearchMatrix(t *testing.T) {
	fmt.Println(searchMatrix([][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 20))
}
