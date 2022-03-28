package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) < 1 || len(inorder) < 1 {
		return nil
	}
	// 寻找中序数组中的根节点下标
	root := findRoot(preorder[0], inorder)
	// 递归构建二叉树
	treeNode := &TreeNode{
		Val: preorder[0],
		// 左子树就是前序数组的[1:root+1]， 中序数组的[0:root]
		Left: buildTree(preorder[1:root+1], inorder[0:root]),
		// 右子树就是前序数组的[root+1:]， 中序数组的[root+1:]
		Right: buildTree(preorder[root+1:], inorder[root+1:]),
	}
	return treeNode
}

// 根据根节点的值寻找根节点在中序数组中的下标
func findRoot(value int, arr []int) int {
	for i, v := range arr {
		if v == value {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(*buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}))
}
