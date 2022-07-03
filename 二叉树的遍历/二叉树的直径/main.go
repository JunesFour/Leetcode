package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	maxLength := 0
	var measure func(root *TreeNode) int
	measure = func(root *TreeNode) int {
		if root == nil {
			return 0
		} else {
			lHeight := measure(root.Left)
			rHeight := measure(root.Right)
			maxLength = max(maxLength, lHeight+rHeight)
			return 1 + max(lHeight, rHeight)
		}
	}
	measure(root)
	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {

}
