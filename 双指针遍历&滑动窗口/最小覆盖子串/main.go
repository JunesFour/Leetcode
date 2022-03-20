package main

import "fmt"

func minWindow(s string, t string) string {
	// tArr统计t串中字符的个数， sArr统计s串中字符的个数
	sArr, tArr := make([]int, 128), make([]int, 128)
	slen, tlen := len(s), len(t)
	// 首先用tArr数组统计t串的字符个数
	for i := 0; i < tlen; i++ {
		tArr[t[i]-'A']++
	}
	// left，right分别表示滑动窗口左边界和右边界的下标
	// length表示滑动窗口的大小
	left, right, length := 0, -1, slen+1
	// 遍历s数组，同时统计s串的字符个数，看看有没有将t串覆盖
	// 如果s串覆盖了t串，就尝试着减小窗口的大小
	for i, j := 0, 0; j < slen; j++ {
		sArr[s[j]-'A']++
		// 在s串覆盖t串的条件下，减小窗口的大小
		for equal(sArr, tArr) {
			// 更新窗口大小
			if length > j-i+1 {
				length = j - i + 1
				left = i
				right = j
			}
			// 尝试着剔除窗口左边界的元素
			sArr[s[i]-'A']--
			i++
		}
	}
	return s[left : right+1]
}

// 判断s串有没有覆盖t串
func equal(arr1 []int, arr2 []int) bool {
	for i := 0; i < 128; i++ {
		if arr1[i] < arr2[i] && arr2[i] != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
}
