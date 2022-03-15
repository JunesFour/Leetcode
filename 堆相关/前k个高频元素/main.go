package main

import "fmt"

func topKFrequent(nums []int, k int) []int {
	hashTable := make(map[int]int)
	for _, v := range nums {
		hashTable[v]++
	}
	var count [][]int
	for k, v := range hashTable {
		count = append(count, []int{k, v})
	}
	res := make([]int, 0)
	// 建堆
	for i := len(count)/2 - 1; i >= 0; i-- {
		heapify(count, len(count), i)
	}
	for i := len(count) - 1; i >= 0; i-- {
		count[i], count[0] = count[0], count[i]
		heapify(count, i, 0)
		res = append(res, count[i][0])
		if len(count)-i == k {
			break
		}
	}
	return res
}

func heapify(arr [][]int, n int, i int) {
	//堆的维护
	largest := i    //i为当前父节点的下标
	lson := i*2 + 1 //左孩子
	rson := i*2 + 2 //右孩子
	// 保存最大值的下标
	if lson < n && arr[largest][1] < arr[lson][1] {
		largest = lson
	}
	if rson < n && arr[largest][1] < arr[rson][1] {
		largest = rson
	}
	if largest != i {
		arr[largest], arr[i] = arr[i], arr[largest]
		// 往孩子节点递归维护堆
		heapify(arr, n, largest)
	}
}

func main() {
	fmt.Println(topKFrequent([]int{1, 2}, 2))
}
