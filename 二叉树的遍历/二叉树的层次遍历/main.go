package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder1(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	// 创建一个q队列，每次将一层的节点都入队
	// 第一次将根节点入队
	q := []*TreeNode{root}
	// 不停的入队一层的节点
	for i := 0; len(q) > 0; i++ {
		// 新开一层
		res = append(res, []int{})
		// p队列用来将当前层所有节点的孩子节点入队，即入队下一层的元素
		var p []*TreeNode
		// 遍历当前层节点
		for j := 0; j < len(q); j++ {
			// 将当前层的所有节点添加到res数组中
			node := q[j]
			res[i] = append(res[i], node.Val)
			// 入队下一层节点
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		// 遍历下一层
		q = p
	}
	return res
}

func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	var dfs func(node *TreeNode, depth int)
	dfs = func(node *TreeNode, depth int) {
		// 递归出口
		if node == nil {
			return
		}
		// 初始化新的一层的储存空间
		if len(res) == depth {
			res = append(res, []int{})
		}
		// 将第depth层的节点添加到res中
		res[depth] = append(res[depth], node.Val)
		if node.Left != nil {
			dfs(node.Left, depth+1)
		}
		if node.Right != nil {
			dfs(node.Right, depth+1)
		}
	}
	dfs(root, 0)
	return res
}

func main() {

}
