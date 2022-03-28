package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) < 1 || len(postorder) < 1 {
		return nil
	}
	// 根节点就是后续数组的最后一位
	root := findRoot(inorder, postorder[len(postorder)-1])
	return &TreeNode{
		Val: postorder[len(postorder)-1],
		// 左子树就是中序数组的inorder[:root], 后续数组的postorder[:root]
		Left: buildTree(inorder[:root], postorder[:root]),
		// 右子树就是中序数组的inorder[root+1:]，后续数组的postorder[root:len(postorder) - 1]
		Right: buildTree(inorder[root+1:], postorder[root:len(postorder)-1]),
	}
}

func findRoot(arr []int, target int) int {
	for i, v := range arr {
		if v == target {
			return i
		}
	}
	return -1
}

func main() {

}
