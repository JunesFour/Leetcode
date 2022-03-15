package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	sortForArray(intervals)
	var result = make([][]int, 0)
	var left, right = intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		var head, rear = intervals[i][0], intervals[i][1]
		if head > right {
			result = append(result, []int{left, right})
			left = head
			right = rear
		} else if head >= left && rear >= right {
			right = rear
		}
	}
	result = append(result, []int{left, right})
	return result
}
func sortForArray(intervals [][]int) {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		} else {
			return intervals[i][0] < intervals[j][0]
		}
	})
}

func main() {
	/*intervals := [][]int{
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	}*/

	intervals := [][]int{
		{2, 6},
		{1, 3},

		{8, 10},
		{15, 18},
	}

	//fmt.Println(merge(intervals))
	sortForArray(intervals)
	fmt.Println(intervals)
}
