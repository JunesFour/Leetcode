package main

import "fmt"

var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
var ans int

func getMaximumGold(grid [][]int) int {
	fmt.Println(ans)
	var dfs func(x, y, cnt int)
	dfs = func(x, y, cnt int) {
		cnt += grid[x][y]
		if cnt > ans {
			ans = cnt
		}
		res := grid[x][y]
		grid[x][y] = 0
		for _, v := range dirs {
			nowX := x + v.x
			nowY := y + v.y
			if nowX >= 0 && nowX < len(grid) && nowY >= 0 && nowY < len(grid[0]) && grid[nowX][nowY] > 0 {
				dfs(nowX, nowY, cnt)
			}
		}
		grid[x][y] = res
	}

	for i, v := range grid {
		for j, vv := range v {
			if vv > 0 {
				dfs(i, j, 0)
			}
		}
	}

	return ans

}

func main() {
	getMaximumGold([][]int{})
}
