package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// root.Val = root.Val + root.Right.Val
// root.Left.Val = root.Val
// 反向中序遍历即可
func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	var traversal func(root *TreeNode)
	traversal = func(root *TreeNode) {
		if root != nil {
			traversal(root.Right)
			sum += root.Val
			root.Val = sum
			traversal(root.Left)
		}
	}
	traversal(root)
	return root
}

func main() {

}
