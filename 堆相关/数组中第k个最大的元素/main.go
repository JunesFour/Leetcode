package main

import "fmt"

func findKthLargest(nums []int, k int) int {
	// 建堆
	for i := len(nums)/2 - 1; i >= 0; i-- {
		heapify(nums, len(nums), i)
	}
	for i := len(nums) - 1; i > 0; i-- {
		nums[i], nums[0] = nums[0], nums[i]
		heapify(nums, i, 0)
		if len(nums)-i == k {
			return nums[i]
		}
	}
	return nums[0]
}

func heapify(arr []int, n int, i int) {
	//堆的维护
	largest := i    //i为当前父节点的下标
	lson := i*2 + 1 //左孩子
	rson := i*2 + 2 //右孩子
	// 保存最大值的下标
	if lson < n && arr[largest] < arr[lson] {
		largest = lson
	}
	if rson < n && arr[largest] < arr[rson] {
		largest = rson
	}
	if largest != i {
		arr[largest], arr[i] = arr[i], arr[largest]
		// 往孩子节点递归维护堆
		heapify(arr, n, largest)
	}
}

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
}
