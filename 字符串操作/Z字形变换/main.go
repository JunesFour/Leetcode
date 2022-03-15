package main

import "strings"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	// 二维数组。把Z型字符串同一行的字符存到同一行中
	rows := make([]string, numRows)
	// 计算V型字符串的周期
	n := 2*numRows - 2
	for i, char := range s {
		// 某个字符的下标应该是min(x, n - x)
		x := i % n
		rows[min(x, n-x)] += string(char)
	}
	// 最后将每一行的字符按顺序拼接为一个字符串
	return strings.Join(rows, "")
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {

}
