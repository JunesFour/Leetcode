package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	res := Preorder(root)
	for i := 1; i < len(res); i++ {
		pre, cur := res[i-1], res[i]
		pre.Left, pre.Right = nil, cur
	}
}
func Preorder(root *TreeNode) (res []*TreeNode) {
	if root != nil {
		res = append(res, root)
		res = append(res, Preorder(root.Left)...)
		res = append(res, Preorder(root.Right)...)
	}
	return
}

func flatten1(root *TreeNode) {
	cur := root
	for cur != nil {
		// 首先判断当前节点的左子树是否为空
		if cur.Left != nil {
			next := cur.Left
			// 寻找左子树中的最后一个节点
			// 按照前序遍历的规律，这个节点就是右子树的前驱节点
			pre := cur.Left
			for pre.Right != nil {
				pre = pre.Right
			}
			pre.Right = cur.Right
			cur.Left, cur.Right = nil, next
		}
		cur = cur.Right
	}
}

func main() {
}
