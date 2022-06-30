package main

import (
	"fmt"
	"sort"
)

type lessSort []int

func (arr lessSort) Len() int {
	return len(arr)
}
func (arr lessSort) Less(i, j int) bool {
	return arr[i] > arr[j]
}

func (arr lessSort) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func reconstructQueue(people [][]int) (res [][]int) {
	// 根据一维逆序，二维正序
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]
	})
	res = make([][]int, 0)
	for _, v := range people {
		// 添加新元素到尾部
		res = append(res, v)
		// 把待插入位置的后面所有元素向后移动一位
		copy(res[v[1]+1:], res[v[1]:])
		// 将待插入元素插入到相应位置
		res[v[1]] = v
	}
	return res
}

func main() {
	fmt.Println()
}
