package main

func exist(board [][]byte, word string) bool {
	type pair struct {
		x, y int
	}
	direction := []pair{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	lenX, lenY := len(board), len(board[0])
	visit := make([][]bool, lenX)
	for i := 0; i < lenX; i++ {
		visit[i] = make([]bool, lenY)
	}

	var check func(i, j, k int) bool
	check = func(i, j, k int) bool {
		// 剪枝
		if board[i][j] != word[k] {
			return false
		}
		// 如果找到了所有word中的字符，直接返回true
		if k == len(word)-1 {
			return true
		}
		// 标记已经走过这条路
		visit[i][j] = true
		// 回溯
		defer func() {
			visit[i][j] = false
		}()
		// 遍历四个方向
		for _, dir := range direction {
			newX, newY := i+dir.x, j+dir.y
			if newX >= 0 && newX < lenX && newY >= 0 && newY < lenY && !visit[newX][newY] {
				if check(newX, newY, k+1) {
					return true
				}
			}
		}
		return false
	}
	for i, row := range board {
		for j, _ := range row {
			if check(i, j, 0) {
				return true
			}
		}
	}
	return false
}

func main() {

}
