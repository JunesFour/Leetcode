package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32
	var maxPrice func(node *TreeNode) int
	maxPrice = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		// 计算左子树的最大贡献值
		leftPrice := max(maxPrice(node.Left), 0)
		// 计算右子树的最大贡献值
		rightPrice := max(maxPrice(node.Right), 0)
		// 加上当前节点的路径和
		currPrice := node.Val + leftPrice + rightPrice
		// 更新最大路径和
		maxSum = max(maxSum, currPrice)
		// 返回当前节点的最大贡献值
		return node.Val + max(leftPrice, rightPrice)
	}
	// 从根节点开始递归
	maxPrice(root)
	return maxSum
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
