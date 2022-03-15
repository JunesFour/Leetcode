package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return dfs(root, math.MinInt64, math.MaxInt64)
}

func dfs(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	// 判断是否在这个区间内
	if root.Val <= min || root.Val >= max {
		return false
	}
	// 更新左递归的最大值为当前节点
	// 更新右递归的最小值为当前节点
	return dfs(root.Left, min, root.Val) && dfs(root.Right, root.Val, max)
}

func main() {

}
