package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	result := robTree(root)
	return max(result[0], result[1])
}
func robTree(cur *TreeNode) []int {
	if cur == nil {
		return []int{0, 0}
	}
	// 后续遍历得到左右子树的结果
	left := robTree(cur.Left)
	right := robTree(cur.Right)
	// 偷当前节点
	val1 := cur.Val + left[0] + right[0]
	// 不偷当前节点
	val2 := max(left[0], left[1]) + max(right[0], right[1])
	return []int{val2, val1}
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
