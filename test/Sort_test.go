package main

import "testing"

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

func TestSort(t *testing.T) {

}
