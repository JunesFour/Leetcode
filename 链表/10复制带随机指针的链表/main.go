package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

var cachedNode map[*Node]*Node

func deepCopy(node *Node) *Node {
	// 递归出口
	// 节点为空
	if node == nil {
		return nil
	}
	// 节点存在
	if n, has := cachedNode[node]; has {
		return n
	}
	// 不存在就创建节点
	newNode := &Node{Val: node.Val}
	cachedNode[node] = newNode
	// 递归复制后继节点和随机节点
	newNode.Next = deepCopy(node.Next)
	newNode.Random = deepCopy(node.Random)
	return newNode
}

func copyRandomList(head *Node) *Node {
	cachedNode = map[*Node]*Node{}
	return deepCopy(head)
}

func main() {

}
