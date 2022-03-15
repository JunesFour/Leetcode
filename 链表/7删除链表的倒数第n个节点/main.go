package main

import "fmt"

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
	L := &ListNode{}
	head := L
	// 1 2 3 4 5 6
	for i := 1; i < 7; i++ {
		/*var x int
		fmt.Scanf("%d", &x)*/
		L.Next = &ListNode{Val: i}
		L = L.Next
	}
	head = removeNthFromEnd(head.Next, 3)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
