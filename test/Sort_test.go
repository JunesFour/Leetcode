package main

import (
	"fmt"
	"sort"
	"testing"
)

type sortInt []int

func (arr sortInt) Len() int {
	return len(arr)
}
func (arr sortInt) Less(i, j int) bool {
	return arr[i] < arr[j]
}

func (arr sortInt) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

type Student struct {
	name  string
	score int
}

type Students []Student

func (s Students) Len() int {
	return len(s)
}
func (s Students) Less(i, j int) bool {
	return s[i].score < s[j].score
}
func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func TestSort(t *testing.T) {
	var str []string
	str = append(str, "abcd", "def", "bcd")
	//sort.Strings(str)
	//fmt.Println(str)
	//arr := []int{1, 4, 6, 3, 10}
	//var _arr sortInt = arr
	//sort.Sort(sort.Reverse(_arr))
	//fmt.Println(_arr)
	//sort.Slice(str, func(i, j int) bool {
	//	return str[i] > str[j]
	//})
	//fmt.Println(str)
	//sort.Sort()
	//var students Students
	//students = []Student{
	//	Student{"zhangsan", 89},
	//	Student{"lisi", 98},
	//	Student{"wangwu", 78},
	//}
	//sort.Sort(students)
	//fmt.Println(students)
	students := []Student{
		Student{"zhangsan", 89},
		Student{"lisi", 98},
		Student{"wangwu", 78},
	}
	sort.Slice(students, func(i, j int) bool {
		return students[i].score > students[j].score
	})
	fmt.Println(students)
}
