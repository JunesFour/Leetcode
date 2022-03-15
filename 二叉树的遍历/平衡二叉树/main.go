package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return abs(dfs(root.Left)-dfs(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func dfs(t *TreeNode) int {
	if t == nil {
		return 0
	}
	lChild := dfs(t.Left)
	rChild := dfs(t.Right)
	return 1 + max(lChild, rChild)
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {

}
