package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	var Count func(L *TreeNode) int
	Count = func(L *TreeNode) int {
		if L == nil {
			return 0
		} else {
			lChild := Count(L.Left)
			rChild := Count(L.Right)
			if lChild > rChild {
				return lChild + 1
			} else {
				return rChild + 1
			}
		}
	}
	return Count(root)
}

func main() {

}
