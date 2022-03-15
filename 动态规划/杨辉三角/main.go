package main

import "fmt"

func generate(numRows int) [][]int {
	res := [][]int{{1}}
	if numRows == 1 {
		return res
	}
	res = append(res, []int{1, 1})
	if numRows == 2 {
		return res
	}
	for i := 3; i <= numRows; i++ {
		var rows []int
		for j := 1; j <= i; j++ {
			if j != 1 && j != i {
				rows = append(rows, res[i-2][j-2]+res[i-2][j-1])
			} else {
				rows = append(rows, 1)
			}
		}
		res = append(res, rows)
	}
	return res
}

func main() {
	fmt.Println(generate(5))
}
