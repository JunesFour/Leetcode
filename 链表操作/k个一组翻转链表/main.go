package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 将传入的一段链表翻转
func reverse(head, tail *ListNode) (*ListNode, *ListNode) {
	// 现在的首节点就是后面的尾节点，提前将prev指向尾结点的下一位
	prev := tail.Next
	cur := head
	for prev != tail {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	// prev现在指向尾节点，就是现在的首节点
	// head节点变成尾节点
	return prev, head
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 1 {
		return head
	}
	// 创建一个哑节点，为了避免
	prev := &ListNode{0, head}
	pre := prev
	// 截取一段链表进行翻转
	left, right := head, head
	// 初始化尾指针
	for i := 0; i < k-1; i++ {
		right = right.Next
	}
	count := 0
	for right != nil {
		// 选取一段翻转
		if count%k == 0 {
			left, right = reverse(left, right)
			pre.Next = left
		}
		pre = pre.Next
		left = left.Next
		right = right.Next
		count++
	}
	return prev.Next
}

func main() {

}
