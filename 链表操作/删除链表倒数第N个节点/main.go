package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	pre := &ListNode{Val: 0, Next: head}
	one, two := pre, head
	for i := 0; i < n; i++ {
		two = two.Next
	}
	for two != nil {
		one = one.Next
		two = two.Next
	}
	one.Next = one.Next.Next
	return pre.Next
}

func main() {

}
