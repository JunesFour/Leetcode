package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
		// 往左子树找
	} else if key < root.Val {
		root.Left = deleteNode(root.Left, key)
		// 往右子树找
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
		// 当前节点的值是key
	} else {
		// 无左右子树，直接删除
		if root.Left == nil && root.Right == nil {
			root = nil
			// 有右子树，将后继节点换到当前位置，然后删除后继节点
		} else if root.Right != nil {
			// 替换
			root.Val = Successor(root).Val
			// 删除后继节点
			root.Right = deleteNode(root.Right, root.Val)
			// 有左子树，将前趋节点换到当前位置，然后删除前趋节点
		} else {
			root.Val = Predecessor(root).Val
			root.Left = deleteNode(root.Left, root.Val)
		}
	}
	return root
}

// Predecessor 寻找前趋节点
func Predecessor(node *TreeNode) *TreeNode {
	pre := node.Left
	for pre.Right != nil {
		pre = pre.Right
	}
	return pre
}

// Successor 寻找后继节点
func Successor(node *TreeNode) *TreeNode {
	post := node.Right
	for post.Left != nil {
		post = post.Left
	}
	return post
}

func main() {

}
