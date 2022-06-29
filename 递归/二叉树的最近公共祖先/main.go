package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// root为p或q中的一个
	if root == p || root == q {
		return root
	}
	// 递归遍历左子树和右子树
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	// 如果分别在左右子树中，公共祖先就是root
	if left != nil && right != nil {
		return root
		// 都分布在左子树中
	} else if left != nil {
		return left
		// 都分布在右子树中
	} else {
		return right
	}
}

func main() {

}
