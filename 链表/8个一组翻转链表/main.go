package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverse(head, tail *ListNode) (*ListNode, *ListNode) {
	prev := tail.Next
	cur := head
	for prev != tail {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev, head
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 1 {
		return head
	}
	prev := &ListNode{0, head}
	pre := prev
	left, right := head, head
	for i := 0; i < k-1; i++ {
		right = right.Next
	}
	count := 0
	for right != nil {
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
