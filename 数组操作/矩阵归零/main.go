package main

type Point struct {
	x int
	y int
}

func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	hashTable := make(map[Point]bool)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 && hashTable[Point{i, j}] == false {
				hashTable[Point{i, j}] = true
			}
		}
	}
	for key, _ := range hashTable {
		// 同列变为0
		for i := 0; i < m; i++ {
			matrix[i][key.y] = 0
		}
		// 同行变为0
		for i := 0; i < n; i++ {
			matrix[key.x][i] = 0
		}
	}
}

func main() {

}
