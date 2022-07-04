package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	var dfs func(r1, r2 *TreeNode) *TreeNode
	dfs = func(r1, r2 *TreeNode) *TreeNode {
		if r1 == nil {
			return r2
		}
		if r2 == nil {
			return r1
		}
		r1.Val = r1.Val + r2.Val
		r1.Left = dfs(r1.Left, r2.Left)
		r1.Right = dfs(r1.Right, r2.Right)
		return r1
	}
	return dfs(root1, root2)
}

func main() {

}
