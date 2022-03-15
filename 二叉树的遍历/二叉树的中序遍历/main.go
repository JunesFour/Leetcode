package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	var dfs func(root *TreeNode, res *[]int)
	dfs = func(root *TreeNode, res *[]int) {
		if root.Left != nil {
			dfs(root.Left, res)
		}
		*res = append(*res, root.Val)
		if root.Right != nil {
			dfs(root.Right, res)
		}
	}
	dfs(root, &res)
	return res
}

func main() {

}
